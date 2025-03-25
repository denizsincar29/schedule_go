package schedule

import (
	"schedule/modeus"
	"schedule/parsers"
	"time"
)

type Schedule struct {
	Modeus *modeus.Modeus
}

func New(email string, password string) (*Schedule, error) {
	m, err := modeus.New(email, password)
	if err != nil {
		return nil, err
	}
	return &Schedule{
		Modeus: m,
	}, nil
}

func (s *Schedule) Close() {
	s.Modeus.Close()
}

func (s *Schedule) GetSchedule(personID string, startTime time.Time, endTime time.Time) (*parsers.Events, error) {
	sched, err := s.Modeus.GetSchedule(personID, startTime, endTime)
	if err != nil {
		return &parsers.Events{}, err
	}
	return parsers.ParseEvents(sched)
}

// search for a person by name or id
func (s *Schedule) SearchPerson(query string, byID bool) (*parsers.People, error) {
	people, err := s.Modeus.SearchPerson(query, byID)
	if err != nil {
		return &parsers.People{}, err
	}
	return parsers.ParsePeople(people), nil
}
