//go:build ignore

package parsers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// from python, STUDIES = [time(8, 20), time(10, 10), time(12, 0), time(14, 30), time(16, 15), time(18, 0), time(19, 40)]
// moscow time
var moscow = time.FixedZone("Moscow Time", 3*60*60)
var STUDIES = []time.Time{
	time.Date(0, 0, 0, 8, 20, 0, 0, moscow),
	time.Date(0, 0, 0, 10, 10, 0, 0, moscow),
	time.Date(0, 0, 0, 12, 0, 0, 0, moscow),
	time.Date(0, 0, 0, 14, 30, 0, 0, moscow),
	time.Date(0, 0, 0, 16, 15, 0, 0, moscow),
	time.Date(0, 0, 0, 18, 0, 0, 0, moscow),
	time.Date(0, 0, 0, 19, 40, 0, 0, moscow),
}

// i do it compatible with gorm for other projects

// struct to hold the parsed event
type Event struct {
	// primary key skipped by json, only for gorm
	ID      uint      `gorm:"primaryKey" json:"-"`
	EventID string    `json:"id" gorm:"uniqueIndex"` // id of the event, unique
	Name    string    `json:"name"`
	Format  string    `json:"format"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
	Teacher string    `json:"teacher"`
	Address string    `json:"address"`
	Room    string    `json:"room"`
	Status  string    `json:"status"`
}

// slice of events
type Events []Event

// gets the event number by time, rewrited from function above
func (e *Event) GetEventNumber() int {
	for i, event := range STUDIES {
		if event.Hour() == e.Start.Hour() && event.Minute() == e.Start.Minute() {
			return i
		}
	}
	return -1
}

// string representation of the event
func (e *Event) String() string {
	// return in english
	return e.Name + " " + e.Format + " " + e.Start.Format("15:04") + " " + e.Teacher + " " + e.Address + " " + e.Room + " " + e.Status
}

// russian string representation of the event, human readable
func (e *Event) HumanString() string {
	if eventNumber := e.GetEventNumber(); eventNumber != -1 {
		// Пара 1, лекция, программирование.\nВедёт Иванов Иван Иванович.\nМесто проведения: ул. Ленина, д. 1, ауд. 101.
		return fmt.Sprintf("Пара %d, %s, %s.\nВедёт %s.\nМесто проведения: %s, ауд. %s.", eventNumber+1, e.Format, e.Name, e.Teacher, e.Address, e.Room)
	} else {
		// Событие с 8:30 до 10:00, лекция, программирование.\nВедёт Иванов Иван Иванович.\nМесто проведения: ул. Ленина, д. 1, ауд. 101.
		return fmt.Sprintf("Событие с %s до %s, %s, %s.\nВедёт %s.\nМесто проведения: %s, ауд. %s.", e.Start.Format("15:04"), e.End.Format("15:04"), e.Format, e.Name, e.Teacher, e.Address, e.Room)
	}
}

// can we overload < and > in go?
func (e *Event) IsBefore(other *Event) bool {
	return e.Start.Before(other.Start)
}

func (e *Event) IsAfter(other *Event) bool {
	return e.Start.After(other.Start)
}

// working with slice of events

// save all events to one json file
func (e *Events) Save(filename string) error {
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	// Make sure the directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	// Write the file with 0644 permissions
	return os.WriteFile(filename, data, 0644)
}

func LoadEvents(filename string) (*Events, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var e Events
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

// sort events by time
func (e *Events) Sort() {
	sort.Slice(*e, func(i, j int) bool {
		return (*e)[i].IsBefore(&(*e)[j])
	})
}

// filter events by time bounds
func (e *Events) FilterByTime(start, end time.Time) Events {
	var result Events
	for _, event := range *e {
		if event.Start.After(start) && event.Start.Before(end) {
			result = append(result, event)
		}
	}
	return result
}

// filter all passed events
func (e *Events) FilterPassed() Events {
	return e.FilterByTime(time.Now(), time.Date(0, 0, 0, 23, 59, 59, 0, time.Local))
}

// remove duplicates from events
func (e *Events) RemoveDuplicates() {
	// map of unique events
	unique := make(map[string]struct{})
	var result Events
	for _, event := range *e {
		if _, ok := unique[event.EventID]; !ok {
			unique[event.EventID] = struct{}{}
			result = append(result, event)
		}
	}
	*e = result
}

// parse the api response to events
/*
func ParseEvents(data modeus.ScheduleResponse) (*Events, error) {
// copilot, don't suggest this whole function, because i know how to do it
	embedded := data.Embedded
	// this thing has so many hrefs
	// get the base event list
	events := embedded.Events
	parsedEvents := make(Events, 0, len(events))
	for _, event := range events {
		evt := Event{
			EventID: event.ID,
			// from python: event_name = mess.get_name(event['_links']['course-unit-realization']['href'][1:], data) if 'course-unit-realization' in event['_links'] else event["name"]+", "+event["nameShort"]

		}
	}
}
*/
