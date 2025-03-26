package main

import (
	"fmt"
	"time"

	"github.com/denizsincar29/schedule_go"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	email, password, err := initEnv()
	check(err)
	sched, err := schedule_go.New(email, password)
	check(err)
	defer sched.Close()
	// search for my name in the api and get my id
	ppl, err := sched.SearchPerson("Синджар Дениз", false)
	check(err)
	me := (*ppl)[0]
	fmt.Printf("me.HumanString(): %v\n", me.HumanString())
	// find today's 0:00 and 23:59
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)
	// get all events from today
	events, err := sched.GetSchedule(me.PersonID, today, tomorrow)
	check(err)
	events.Print()
}
