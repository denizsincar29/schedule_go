package parsers

import (
	"fmt"

	"time"

	"github.com/denizsincar29/schedule_go/modeus"
)

// Constants for the type of person (student or teacher)
type Role int

const (
	Student Role = iota
	Teacher
)

// stringify
func (r Role) String() string {
	return [...]string{"student", "teacher"}[r]
}

// jsonify
func (r Role) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}

// Person struct
type Person struct {
	PersonID  string    // from api
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Specialty string    `json:"specialty,omitempty"`
	Profile   string    `json:"profile,omitempty"`
	Group     string    `json:"group,omitempty"`
	Role      Role      `json:"role"`
}

type People []Person

// parse people from api response
func ParsePeople(data *modeus.SearchPersonResponse) *People {
	var people People = make([]Person, 0, len(data.Embedded.Persons))
	for _, apiPerson := range data.Embedded.Persons {
		startDate, endDate := data.GetPersonDates(&apiPerson)
		//info1, info2 = data.GetPersonInfo(&apiPerson)  // specialti/profile or group / empty string
		specialty := ""
		profile := ""
		group := ""
		apiRole := data.GetPersonType(&apiPerson)
		role := Role(apiRole)
		if apiRole == modeus.RoleStudent {
			specialty, profile = data.GetPersonInfo(&apiPerson)
		} else {
			group, _ = data.GetPersonInfo(&apiPerson)
		}
		person := Person{
			PersonID:  apiPerson.ID,
			Name:      apiPerson.FullName, // last name, firstname and russian father name
			StartDate: startDate,
			EndDate:   endDate,
			Specialty: specialty,
			Profile:   profile,
			Group:     group,
			Role:      role,
		}
		people = append(people, person)
	}
	return &people
}

// stringify person
func (p Person) String() string {
	return fmt.Sprintf("PersonID: %s, Name: %s, StartDate: %s, EndDate: %s, Specialty: %s, Profile: %s, Group: %s, Role: %s", p.PersonID, p.Name, p.StartDate, p.EndDate, p.Specialty, p.Profile, p.Group, p.Role)
}

// human readable russian string
func (p Person) HumanString() string {
	// if student:
	// Студент Иванов Иван Иванович, specialty, profile
	// if teacher:
	// Преподаватель Иванов Иван Иванович, group
	if p.Role == Student {
		return fmt.Sprintf("Студент %s, %s, %s", p.Name, p.Specialty, p.Profile)
	} else {
		return fmt.Sprintf("Преподаватель %s, %s", p.Name, p.Group)
	}
}
