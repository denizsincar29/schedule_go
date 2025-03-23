package main

import (
	"os"
	"schedule/modeus"
	"schedule/parsers"
	"time"

	"github.com/joho/godotenv"
)

// Modeus api i allowes us to get schedule not only for me but for any student
// so we will need to provide student id along with email and password

func initEnv() (string, string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", "", err
	}
	email := os.Getenv("MODEUS_EMAIL")
	password := os.Getenv("MODEUS_PASSWORD")
	return email, password, nil
}

func main() {
	email, password, err := initEnv()
	if err != nil {
		panic(err)
	}
	m, err := modeus.New(email, password)
	if err != nil {
		panic(err)
	}
	defer m.Close()
	response, err := m.SearchPerson("Синджар", false) // by id = false
	if err != nil {
		panic(err)
	}
	personList := response.Embedded.Persons
	if len(personList) == 0 {
		println("No person found")
		return
	}
	me := personList[0]
	// tomorrow 00:00 - 23:59
	tomorrow := time.Now().AddDate(0, 0, 1).Truncate(24 * time.Hour)
	tomorrowEnd := tomorrow.Add(24 * time.Hour).Add(-1 * time.Second)
	schedule, err := m.GetSchedule(me.ID, tomorrow, tomorrowEnd)
	if err != nil {
		panic(err)
	}
	events, err := parsers.ParseEvents(schedule)
	if err != nil {
		panic(err)
	}
	events.Print()
}
