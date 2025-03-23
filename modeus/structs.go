package modeus

import "time"

type Link struct {
	Href string `json:"href"`
}

// ScheduleResponse represents the entire schedule response
type ScheduleResponse struct {
	Embedded struct {
		Events                 []Event                 `json:"events"`
		CourseUnitRealizations []CourseUnitRealization `json:"course-unit-realizations"`
		CycleRealizations      []CycleRealization      `json:"cycle-realizations"`
		LessonRealizationTeams []LessonRealizationTeam `json:"lesson-realization-teams"`
		LessonRealizations     []LessonRealization     `json:"lesson-realizations"`
		EventLocations         []EventLocation         `json:"event-locations"`
		Durations              []Duration              `json:"durations"`
		EventRooms             []EventRoom             `json:"event-rooms"`
		Rooms                  []Room                  `json:"rooms"`
		Buildings              []Building              `json:"buildings"`
		EventTeams             []EventTeam             `json:"event-teams"`
		EventOrganizers        []EventOrganizer        `json:"event-organizers"`
		EventAttendees         []EventAttendee         `json:"event-attendees"`
		Persons                []Person                `json:"persons"`
	} `json:"_embedded"`
	Page PageInfo `json:"page"`
}

type PageInfo struct {
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Number        int `json:"number"`
}

type Event struct {
	Name                      string        `json:"name"`
	NameShort                 string        `json:"nameShort"`
	Description               interface{}   `json:"description"`
	TypeID                    string        `json:"typeId"`
	FormatID                  interface{}   `json:"formatId"`
	Start                     time.Time     `json:"start"`
	End                       time.Time     `json:"end"`
	StartsAtLocal             string        `json:"startsAtLocal"`
	EndsAtLocal               string        `json:"endsAtLocal"`
	StartsAt                  string        `json:"startsAt"`
	EndsAt                    string        `json:"endsAt"`
	HoldingStatus             HoldingStatus `json:"holdingStatus"`
	RepeatedLessonRealization interface{}   `json:"repeatedLessonRealization"`
	UserRoleIds               []string      `json:"userRoleIds"`
	LessonTemplateID          string        `json:"lessonTemplateId"`
	Version                   int           `json:"__version"`
	Links                     EventLinks    `json:"_links"`
	ID                        string        `json:"id"`
}

type HoldingStatus struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	AudModifiedAt       string `json:"audModifiedAt"`
	AudModifiedBy       string `json:"audModifiedBy"`
	AudModifiedBySystem bool   `json:"audModifiedBySystem"`
}

type EventLinks struct {
	Self                      Link `json:"self"`
	Type                      Link `json:"type"`
	TimeZone                  Link `json:"time-zone"`
	Grid                      Link `json:"grid"`
	CourseUnitRealization     Link `json:"course-unit-realization"`
	CycleRealization          Link `json:"cycle-realization"`
	LessonRealization         Link `json:"lesson-realization"`
	LessonRealizationTeam     Link `json:"lesson-realization-team"`
	LessonRealizationTemplate Link `json:"lesson-realization-template"`
	HoldingStatusModifiedBy   Link `json:"holding-status-modified-by"`
	Location                  Link `json:"location"`
	Duration                  Link `json:"duration"`
	Team                      Link `json:"team"`
	Organizers                Link `json:"organizers"`
}

type CourseUnitRealization struct {
	Name        string                     `json:"name"`
	NameShort   string                     `json:"nameShort"`
	PrototypeID string                     `json:"prototypeId"`
	Links       CourseUnitRealizationLinks `json:"_links"`
	ID          string                     `json:"id"`
}

type CourseUnitRealizationLinks struct {
	Self           Link `json:"self"`
	PlanningPeriod Link `json:"planning-period"`
}

type CycleRealization struct {
	Name                           string                `json:"name"`
	NameShort                      string                `json:"nameShort"`
	Code                           string                `json:"code"`
	CourseUnitRealizationNameShort string                `json:"courseUnitRealizationNameShort"`
	Links                          CycleRealizationLinks `json:"_links"`
	ID                             string                `json:"id"`
}

type CycleRealizationLinks struct {
	Self                  Link `json:"self"`
	CourseUnitRealization Link `json:"course-unit-realization"`
}

type LessonRealizationTeam struct {
	Name               string                     `json:"name"`
	CycleRealizationID string                     `json:"cycleRealizationId"`
	Links              LessonRealizationTeamLinks `json:"_links"`
	ID                 string                     `json:"id"`
}

type LessonRealizationTeamLinks struct {
	Self Link `json:"self"`
}

type LessonRealization struct {
	Name        string                 `json:"name"`
	NameShort   string                 `json:"nameShort"`
	PrototypeID string                 `json:"prototypeId"`
	Ordinal     int                    `json:"ordinal"`
	Links       LessonRealizationLinks `json:"_links"`
	ID          string                 `json:"id"`
}

type LessonRealizationLinks struct {
	Self Link `json:"self"`
}

type EventLocation struct {
	EventID        string             `json:"eventId"`
	CustomLocation interface{}        `json:"customLocation"`
	Links          EventLocationLinks `json:"_links"`
}

type EventLocationLinks struct {
	Self       []Link `json:"self"`
	EventRooms Link   `json:"event-rooms"`
}

type Duration struct {
	EventID    string        `json:"eventId"`
	Value      int           `json:"value"`
	TimeUnitID string        `json:"timeUnitId"`
	Minutes    int           `json:"minutes"`
	Links      DurationLinks `json:"_links"`
}

type DurationLinks struct {
	Self     []Link `json:"self"`
	TimeUnit Link   `json:"time-unit"`
}

type EventRoom struct {
	Links EventRoomLinks `json:"_links"`
	ID    string         `json:"id"`
}

type EventRoomLinks struct {
	Self  Link `json:"self"`
	Event Link `json:"event"`
	Room  Link `json:"room"`
}

type Room struct {
	Name               string      `json:"name"`
	NameShort          string      `json:"nameShort"`
	Building           Building    `json:"building"`
	ProjectorAvailable bool        `json:"projectorAvailable"`
	TotalCapacity      int         `json:"totalCapacity"`
	WorkingCapacity    int         `json:"workingCapacity"`
	DeletedAtUtc       interface{} `json:"deletedAtUtc"`
	Links              RoomLinks   `json:"_links"`
	ID                 string      `json:"id"`
}

type RoomLinks struct {
	Self     Link `json:"self"`
	Type     Link `json:"type"`
	Building Link `json:"building"`
}

type Building struct {
	Name              string        `json:"name"`
	NameShort         string        `json:"nameShort"`
	Address           string        `json:"address"`
	SearchableAddress interface{}   `json:"searchableAddress"`
	DisplayOrder      int           `json:"displayOrder"`
	Links             BuildingLinks `json:"_links"`
	ID                string        `json:"id"`
}

type BuildingLinks struct {
	Self Link `json:"self"`
}

type EventTeam struct {
	EventID string         `json:"eventId"`
	Size    int            `json:"size"`
	Links   EventTeamLinks `json:"_links"`
}

type EventTeamLinks struct {
	Self  Link `json:"self"`
	Event Link `json:"event"`
}

type EventOrganizer struct {
	EventID string              `json:"eventId"`
	Links   EventOrganizerLinks `json:"_links"`
}

type EventOrganizerLinks struct {
	Self           Link `json:"self"`
	Event          Link `json:"event"`
	EventAttendees Link `json:"event-attendees"`
}

type EventAttendee struct {
	RoleID           string             `json:"roleId"`
	RoleName         string             `json:"roleName"`
	RoleNamePlural   string             `json:"roleNamePlural"`
	RoleDisplayOrder int                `json:"roleDisplayOrder"`
	Links            EventAttendeeLinks `json:"_links"`
	ID               string             `json:"id"`
}

type EventAttendeeLinks struct {
	Self   Link `json:"self"`
	Event  Link `json:"event"`
	Person Link `json:"person"`
}

type Person struct {
	LastName   string      `json:"lastName"`
	FirstName  string      `json:"firstName"`
	MiddleName string      `json:"middleName"`
	FullName   string      `json:"fullName"`
	Links      PersonLinks `json:"_links"`
	ID         string      `json:"id"`
}

type PersonLinks struct {
	Self Link `json:"self"`
}

// SearchPersonResponse represents the response from a person search
type SearchPersonResponse struct {
	Embedded struct {
		Persons   []Person   `json:"persons"`
		Students  []Student  `json:"students"`
		Employees []Employee `json:"employees"`
	} `json:"_embedded"`
	Page PageInfo `json:"page"`
}

type Student struct {
	ID                string    `json:"id"`
	PersonID          string    `json:"personId"`
	FlowID            string    `json:"flowId"`
	FlowCode          string    `json:"flowCode"`
	SpecialtyCode     string    `json:"specialtyCode"`
	SpecialtyName     string    `json:"specialtyName"`
	SpecialtyProfile  string    `json:"specialtyProfile"`
	LearningStartDate time.Time `json:"learningStartDate"`
	LearningEndDate   time.Time `json:"learningEndDate"`
}

type Employee struct {
	ID        string    `json:"id"`
	PersonID  string    `json:"personId"`
	GroupID   string    `json:"groupId"`
	GroupName string    `json:"groupName"`
	DateIn    time.Time `json:"dateIn"`
	DateOut   time.Time `json:"dateOut"`
}
