package main

import (
	"fmt"
	"os"
	"schedule/modeus"
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
	start := time.Now().Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)
	// get schedule for the day
	// now modeus has function SearchPerson! Lets seach for Синджар Дениз
	personList, err := m.SearchPerson("Синджар Дениз", false) // by id = false
	if err != nil {
		panic(err)
	}

	persid := personList.Embedded.Persons[0].ID // it will be some kinda better after we make a good parser
	result, err := m.GetSchedule(persid, start, end)
	if err != nil {
		panic(err)
	}
	for _, event := range result.Embedded.Events {
		fmt.Printf("event.Name: %v\n", event.Name)
		fmt.Printf("event.Start: %v\n", event.Start)
		fmt.Printf("event.End: %v\n", event.End)
	}

}
