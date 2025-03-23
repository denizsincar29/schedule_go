package schedule

import (
	"schedule/modeus"
	"schedule/parsers"
	"time"
)

type Schedule struct {
	Modeus modeus.Modeus
}

func NewSchedule(email string, password string) (*Schedule, error) {
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
