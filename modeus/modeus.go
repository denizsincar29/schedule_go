package modeus

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	defaultBaseURL = "https://narfu.modeus.org"
	defaultAuthURL = "https://narfu-auth.modeus.org"
	tokenLifetime  = 12 * time.Hour
)

// a helper function that takes ref to http request and if wasm is used, sets the header
func setWasm(req *http.Request) {
	// if wasm is used, set the header
	if runtime.GOARCH == "wasm" {
		log.Println("WASM detected, setting headers")
		req.Header.Set("js.fetch:mode", "no-cors")
		req.Header.Set("js.fetch:credentials", "include")

	}
}

type Modeus struct {
	httpClient *http.Client
	baseURL    string
	authURL    string
	email      string
	password   string
	token      string
	expiry     time.Time
	mu         sync.RWMutex
}

func New(email, password string) (*Modeus, error) {
	return NewWithConfig(defaultBaseURL, defaultAuthURL, email, password)
}

func NewWithConfig(baseURL, authURL string, email, password string) (*Modeus, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create cookie jar: %w", err)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		Jar:     jar,
	}

	modeus := &Modeus{
		httpClient: client,
		baseURL:    baseURL,
		authURL:    authURL,
		email:      email,
		password:   password,
	}

	if err := modeus.checkToken(); err != nil {
		return nil, fmt.Errorf("failed to initialize Modeus with valid token: %w", err)
	}
	return modeus, nil
}

// close the client
func (m *Modeus) Close() {
	m.httpClient.CloseIdleConnections()
}

// parses the api token from the login page
func (m *Modeus) parseToken() (string, time.Time, error) {
	email := correctEmail(m.email)
	params := url.Values{
		"client_id":     {"YDNCeCPsf1zL2etGQflijyfzo88a"},
		"redirect_uri":  {"https://narfu.modeus.org/"},
		"response_type": {"id_token"},
		"scope":         {"openid"},
		"state":         {"abab35fcb9164912aa46d287a594a338"},
		"nonce":         {"08cd3a21e9724040acb48cf3a35b0c4b"},
	}
	baseURL := m.authURL + "/oauth2/authorize"
	req, err := http.NewRequest("GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to create request: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	setWasm(req) // WASM header

	beforeForm1Response, err := m.httpClient.Do(req)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to get before form1: %w", err)
	}
	defer beforeForm1Response.Body.Close()
	log.Println("got beforeForm1Response", beforeForm1Response.Request.URL.String())
	form1URL, err := url.Parse(beforeForm1Response.Request.URL.String())
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to parse form1 URL: %w", err)
	}

	req, err = http.NewRequest("GET", form1URL.String(), nil) // Re-create request for wasm headers
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to create request for form1: %w", err)
	}
	req = req.WithContext(ctx)
	setWasm(req) // WASM header

	form1Response, err := m.httpClient.Do(req) // Use the new request with headers.
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to get form1: %w", err)
	}
	defer form1Response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(form1Response.Body)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to parse form1 HTML: %w", err)
	}

	var form1URLMatch string
	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		action, exists := s.Attr("action")
		if exists {
			form1URLMatch = action
			return
		}
	})

	if form1URLMatch == "" {
		return "", time.Time{}, fmt.Errorf("parseToken: can't parse 1st form")
	}

	data := url.Values{
		"UserName":   {email},
		"Password":   {m.password},
		"AuthMethod": {"FormsAuthentication"},
	}

	form2URL, err := url.Parse(form1URLMatch)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to parse form2 URL: %w", err)
	}
	req, err = http.NewRequest("POST", form2URL.String(), nil)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to create request for form2: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(bytes.NewBufferString(data.Encode()))
	req = req.WithContext(ctx)
	setWasm(req)

	form2Response, err := m.httpClient.Do(req) // Use the new request with WASM Headers
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to post to form1: %w", err)
	}
	defer form2Response.Body.Close()

	form2Doc, err := goquery.NewDocumentFromReader(form2Response.Body)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to parse form2 HTML: %w", err)
	}

	form2Matches := make(map[string]string)
	form2Doc.Find("input[type='hidden']").Each(func(i int, s *goquery.Selection) {
		name, existsName := s.Attr("name")
		value, existsValue := s.Attr("value")

		if existsName && existsValue {
			form2Matches[name] = value
		}
	})
	if len(form2Matches) < 2 {
		return "", time.Time{}, fmt.Errorf("parseToken: can't parse 2nd form")
	}

	form2Data := url.Values{}
	for k, v := range form2Matches {
		form2Data.Set(k, v)
	}
	form2Data.Set("UserName", email) //need to resend, if not it returns user name error
	form2Data.Set("Password", m.password)

	url3 := m.authURL + "/commonauth"
	req, err = http.NewRequest("POST", url3, nil)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to create request for commonauth: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Body = io.NopCloser(bytes.NewBufferString(form2Data.Encode()))
	req = req.WithContext(ctx)

	setWasm(req)
	lastResponse, err := m.httpClient.Do(req) // Use the new request with headers

	if err != nil {
		return "", time.Time{}, fmt.Errorf("failed to post to form2: %w", err)
	}
	defer lastResponse.Body.Close()

	lastURL := lastResponse.Request.URL.String()

	idTokenRegex := regexp.MustCompile(`id_token=(.+?)&`)
	idTokenMatch := idTokenRegex.FindStringSubmatch(lastURL)

	if len(idTokenMatch) < 2 {
		return "", time.Time{}, fmt.Errorf("parseToken: can't parse id_token")
	}

	idToken := idTokenMatch[1]

	// Token expires 12 hours after creation.
	expiry := time.Now().Add(tokenLifetime)
	return idToken, expiry, nil
}

// checkToken checks the token's validity and refreshes it if needed.
func (m *Modeus) checkToken() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if time.Until(m.expiry) > tokenLifetime/2 { //half of token life time
		// Token is still valid for more than 6 hours, no need to refresh
		return nil
	}

	newToken, expiry, err := m.parseToken()
	if err != nil {
		return fmt.Errorf("failed to refresh token: %w", err)
	}

	m.token = newToken
	m.expiry = expiry
	return nil
}

// GetToken returns the current valid token. It handles refreshing the token if necessary.
func (m *Modeus) GetToken() (string, error) {
	if err := m.checkToken(); err != nil { // Check and refresh if needed
		return "", err
	}

	m.mu.RLock() // Read-lock for concurrent access
	defer m.mu.RUnlock()
	return m.token, nil
}

// makeRequest sends a request to the Modeus API.
func (m *Modeus) makeRequest(url string, requestJSON map[string]interface{}) ([]byte, error) {
	token, err := m.GetToken() // Get a valid token (refreshes if needed)
	if err != nil {
		return nil, err
	}

	requestBody, err := json.Marshal(requestJSON)
	if err != nil {
		return nil, err
	}

	method := "POST"
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// save body to a file for debugging
	ioutil.WriteFile("response.json", body, 0644)

	return body, nil
}

// GetSchedule returns the schedule for a person with the given ID between the given times.
func (m *Modeus) GetSchedule(personID string, startTime time.Time, endTime time.Time) (*ScheduleResponse, error) {
	url := m.baseURL + "/schedule-calendar-v2/api/calendar/events/search?tz=Europe/Moscow"

	requestJSON := map[string]interface{}{
		"size":             500,
		"timeMin":          startTime.UTC().Format(time.RFC3339),
		"timeMax":          endTime.UTC().Format(time.RFC3339),
		"attendeePersonId": []string{personID},
	}

	response, err := m.makeRequest(url, requestJSON)
	if err != nil {
		return nil, err
	}

	var schedule ScheduleResponse
	err = json.Unmarshal(response, &schedule)
	return &schedule, err

}

// SearchPerson searches for a person by name or ID.
func (m *Modeus) SearchPerson(term string, byID bool) (*SearchPersonResponse, error) {
	url := m.baseURL + "/schedule-calendar-v2/api/people/persons/search"
	mode := "fullName"
	if byID {
		mode = "id"
	}
	requestJSON := map[string]interface{}{
		"size": 10,
		mode:   term,
		"sort": "+fullName",
	}

	response, err := m.makeRequest(url, requestJSON)
	if err != nil {
		return nil, err
	}

	var searchResponse SearchPersonResponse
	err = json.Unmarshal(response, &searchResponse)
	return &searchResponse, err

}

// get who goes on an event
// empty body, url=f"https://narfu.modeus.org/schedule-calendar-v2/api/calendar/events/{event_id}/attendees"
func (m *Modeus) GetEventAttendees(eventID string) (map[string]interface{}, error) {
	url := m.baseURL + "/schedule-calendar-v2/api/calendar/events/" + eventID + "/attendees"
	requestJSON := map[string]interface{}{}
	_, _ = m.makeRequest(url, requestJSON)
	panic("todo")
}
