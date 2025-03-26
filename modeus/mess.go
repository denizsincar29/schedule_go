package modeus

import (
	"strings"
	"time"
)

// mess is because the API response is a huge mess of nested structs

// gets the format and type of an event
func (data *ScheduleResponse) GetFormat(e *Event) string {
	// difficulty level 0
	lessonRealizations := data.Embedded.LessonRealizations
	for _, lr := range lessonRealizations {
		if lr.ID == e.Links.LessonRealization.Href[1:] {
			return strings.TrimSpace(lr.Name + " " + lr.NameShort)
		}
	}
	return "" // can happen in extremely rare cases
}

// Gets the name of an event
func (data *ScheduleResponse) GetEventName(e *Event) string {
	// return e.Name  // if it'd be that easy)))
	//but it's still in difficulty level 1
	for _, cur := range data.Embedded.CourseUnitRealizations {
		if cur.ID == e.Links.CourseUnitRealization.Href[1:] {
			name := cur.Name
			if name == "" {
				return cur.NameShort
			}
			return name
		}
	}
	return "Без названия"
}

// Gets the name of a teacher
func (data *ScheduleResponse) GetTeacherName(e *Event) string {
	// difficulty level 3
	attendees := data.Embedded.EventAttendees
	organizers := data.Embedded.EventOrganizers
	persons := data.Embedded.Persons
	// level 1: find the event_attendee_id
	event_attendee_id := ""
	for _, organizer := range organizers {
		if organizer.EventID == e.ID {
			event_attendee_id = organizer.Links.EventAttendees.Href[1:] // remove the first slash
			break
		}
	}
	// level 2: find the person_id
	person_id := ""
	for _, attendee := range attendees {
		if attendee.ID == event_attendee_id {
			person_id = attendee.Links.Person.Href[1:]
			break
		}
	}
	// level 3: find the person
	for _, person := range persons {
		if person.ID == person_id {
			return person.FullName
		}
	}
	return "Неизвестный преподаватель"
}

// Gets the name of a room
func (data *ScheduleResponse) GetAddress(e *Event) (string, string) {
	// difficulty level 4 because of the nested structure
	locations := data.Embedded.EventLocations
	eventRooms := data.Embedded.EventRooms
	rooms := data.Embedded.Rooms
	// level 1: find the event_room_id
	for _, location := range locations {
		if location.EventID == e.ID {
			// here's where the fun begins
			if location.CustomLocation != nil {
				address, ok := location.CustomLocation.(string)
				if !ok {
					address = "неизвестно"
				}
				room := "" // no room in custom location
				return address, room
			}
			if location.Links.EventRooms.Href == "" {
				return "неизвестно", "неизвестно"
			}

			event_room_id := location.Links.EventRooms.Href[1:]
			// level 2: find the room_id
			for _, event_room := range eventRooms {
				if event_room.ID == event_room_id {
					room_id := event_room.Links.Room.Href[1:]
					// level 3: find the room
					for _, room := range rooms {
						if room.ID == room_id {
							address := strings.Replace(room.Building.Address, "обл. Архангельская, г. Архангельск, ", "", 1)
							return address, room.Name
						}
					}
				}
			}

		}
	}
	return "неизвестно", "неизвестно"
}

// gets the info about a person

type Role int

const (
	RoleStudent Role = iota
	RoleTeacher
	Undefined
)

// gets if a person is student or teacher
func (data *SearchPersonResponse) GetPersonType(person *Person) Role {
	// difficulty level 1
	students := data.Embedded.Students
	teachers := data.Embedded.Employees
	for _, student := range students {
		if student.PersonID == person.ID {
			return RoleStudent
		}
	}
	for _, teacher := range teachers {
		if teacher.PersonID == person.ID {
			return RoleTeacher
		}
	}
	return Undefined // json is terribly wrong and it should never happen! Sorry, api is a mess
}

// gets the specialty and profile of a student or group + nil if it's a teacher
func (data *SearchPersonResponse) GetPersonInfo(person *Person) (string, string) {
	// difficulty level 1
	students := data.Embedded.Students
	teachers := data.Embedded.Employees
	for _, student := range students {
		if student.PersonID == person.ID {
			return student.SpecialtyName, student.SpecialtyProfile
		}
	}
	for _, teacher := range teachers {
		if teacher.PersonID == person.ID {
			return teacher.GroupName, ""
		}
	}
	return "неизвестно", "неизвестно"
}

// gets the date in and date out of the person
func (data *SearchPersonResponse) GetPersonDates(person *Person) (time.Time, time.Time) {
	// difficulty level 1
	students := data.Embedded.Students
	teachers := data.Embedded.Employees
	for _, student := range students {
		if student.PersonID == person.ID {
			return student.LearningStartDate, student.LearningEndDate
		}
	}
	for _, teacher := range teachers {
		if teacher.PersonID == person.ID {
			return teacher.DateIn, teacher.DateOut
		}
	}
	return time.Time{}, time.Time{}
}
