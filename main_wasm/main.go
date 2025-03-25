// go:build wasm
package main

import (
	"log"
	"schedule"
	"syscall/js"
	"time"
)

func LogToDom(msg string) {
	// get the body element
	doc := js.Global().Get("document")
	body := doc.Get("body")
	// create a <p> element
	p := doc.Call("createElement", "p")
	// set the text of the <p> element
	p.Set("textContent", msg)
	// append the <p> element to the body
	body.Call("appendChild", p)

}

// make a log handler that logs to the DOM
func initLogHandler() {
	log.SetFlags(0)
	log.SetOutput(jsWriter{}) // log to the DOM
}

// jsWriter is a type that implements the io.Writer interface
// by writing to the DOM
type jsWriter struct{}

// Write writes p to the DOM
func (jsWriter) Write(p []byte) (n int, err error) {
	LogToDom(string(p))
	return len(p), nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	email, password, err := initEnv()
	check(err)
	sched, err := schedule.New(email, password)
	check(err)
	defer sched.Close()
	// search for my name in the api and get my id
	ppl, err := sched.SearchPerson("Синджар Дениз", false)
	check(err)
	me := (*ppl)[0]
	// find today's 0:00 and 23:59
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.Add(24 * time.Hour)
	// get all events from today
	events, err := sched.GetSchedule(me.PersonID, today, tomorrow)
	check(err)
	// print all events
	for _, event := range *events {
		LogToDom(event.String())
	}
}
