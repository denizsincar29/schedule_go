package modeus

import "time"

type ScheduleResponse struct {
	Embedded struct {
		Events []struct {
			Name          string      `json:"name"`
			NameShort     string      `json:"nameShort"`
			Description   interface{} `json:"description"`
			TypeID        string      `json:"typeId"`
			FormatID      interface{} `json:"formatId"`
			Start         time.Time   `json:"start"`
			End           time.Time   `json:"end"`
			StartsAtLocal string      `json:"startsAtLocal"`
			EndsAtLocal   string      `json:"endsAtLocal"`
			StartsAt      string      `json:"startsAt"`
			EndsAt        string      `json:"endsAt"`
			HoldingStatus struct {
				ID                  string `json:"id"`
				Name                string `json:"name"`
				AudModifiedAt       string `json:"audModifiedAt"`
				AudModifiedBy       string `json:"audModifiedBy"`
				AudModifiedBySystem bool   `json:"audModifiedBySystem"`
			} `json:"holdingStatus"`
			RepeatedLessonRealization interface{} `json:"repeatedLessonRealization"`
			UserRoleIds               []string    `json:"userRoleIds"`
			LessonTemplateID          string      `json:"lessonTemplateId"`
			Version                   int         `json:"__version"`
			Links                     struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Type struct {
					Href string `json:"href"`
				} `json:"type"`
				TimeZone struct {
					Href string `json:"href"`
				} `json:"time-zone"`
				Grid struct {
					Href string `json:"href"`
				} `json:"grid"`
				CourseUnitRealization struct {
					Href string `json:"href"`
				} `json:"course-unit-realization"`
				CycleRealization struct {
					Href string `json:"href"`
				} `json:"cycle-realization"`
				LessonRealization struct {
					Href string `json:"href"`
				} `json:"lesson-realization"`
				LessonRealizationTeam struct {
					Href string `json:"href"`
				} `json:"lesson-realization-team"`
				LessonRealizationTemplate struct {
					Href string `json:"href"`
				} `json:"lesson-realization-template"`
				HoldingStatusModifiedBy struct {
					Href string `json:"href"`
				} `json:"holding-status-modified-by"`
				Location struct {
					Href string `json:"href"`
				} `json:"location"`
				Duration struct {
					Href string `json:"href"`
				} `json:"duration"`
				Team struct {
					Href string `json:"href"`
				} `json:"team"`
				Organizers struct {
					Href string `json:"href"`
				} `json:"organizers"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"events"`
		CourseUnitRealizations []struct {
			Name        string `json:"name"`
			NameShort   string `json:"nameShort"`
			PrototypeID string `json:"prototypeId"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				PlanningPeriod struct {
					Href string `json:"href"`
				} `json:"planning-period"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"course-unit-realizations"`
		CycleRealizations []struct {
			Name                           string `json:"name"`
			NameShort                      string `json:"nameShort"`
			Code                           string `json:"code"`
			CourseUnitRealizationNameShort string `json:"courseUnitRealizationNameShort"`
			Links                          struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				CourseUnitRealization struct {
					Href string `json:"href"`
				} `json:"course-unit-realization"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"cycle-realizations"`
		LessonRealizationTeams []struct {
			Name               string `json:"name"`
			CycleRealizationID string `json:"cycleRealizationId"`
			Links              struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"lesson-realization-teams"`
		LessonRealizations []struct {
			Name        string `json:"name"`
			NameShort   string `json:"nameShort"`
			PrototypeID string `json:"prototypeId"`
			Ordinal     int    `json:"ordinal"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"lesson-realizations"`
		EventLocations []struct {
			EventID        string      `json:"eventId"`
			CustomLocation interface{} `json:"customLocation"`
			Links          struct {
				Self []struct {
					Href string `json:"href"`
				} `json:"self"`
				EventRooms struct {
					Href string `json:"href"`
				} `json:"event-rooms"`
			} `json:"_links"`
		} `json:"event-locations"`
		Durations []struct {
			EventID    string `json:"eventId"`
			Value      int    `json:"value"`
			TimeUnitID string `json:"timeUnitId"`
			Minutes    int    `json:"minutes"`
			Links      struct {
				Self []struct {
					Href string `json:"href"`
				} `json:"self"`
				TimeUnit struct {
					Href string `json:"href"`
				} `json:"time-unit"`
			} `json:"_links"`
		} `json:"durations"`
		EventRooms []struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Event struct {
					Href string `json:"href"`
				} `json:"event"`
				Room struct {
					Href string `json:"href"`
				} `json:"room"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"event-rooms"`
		Rooms []struct {
			Name      string `json:"name"`
			NameShort string `json:"nameShort"`
			Building  struct {
				ID           string `json:"id"`
				Name         string `json:"name"`
				NameShort    string `json:"nameShort"`
				Address      string `json:"address"`
				DisplayOrder int    `json:"displayOrder"`
			} `json:"building"`
			ProjectorAvailable bool        `json:"projectorAvailable"`
			TotalCapacity      int         `json:"totalCapacity"`
			WorkingCapacity    int         `json:"workingCapacity"`
			DeletedAtUtc       interface{} `json:"deletedAtUtc"`
			Links              struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Type struct {
					Href string `json:"href"`
				} `json:"type"`
				Building struct {
					Href string `json:"href"`
				} `json:"building"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"rooms"`
		Buildings []struct {
			Name              string      `json:"name"`
			NameShort         string      `json:"nameShort"`
			Address           string      `json:"address"`
			SearchableAddress interface{} `json:"searchableAddress"`
			DisplayOrder      int         `json:"displayOrder"`
			Links             struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"buildings"`
		EventTeams []struct {
			EventID string `json:"eventId"`
			Size    int    `json:"size"`
			Links   struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Event struct {
					Href string `json:"href"`
				} `json:"event"`
			} `json:"_links"`
		} `json:"event-teams"`
		EventOrganizers []struct {
			EventID string `json:"eventId"`
			Links   struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Event struct {
					Href string `json:"href"`
				} `json:"event"`
				EventAttendees struct {
					Href string `json:"href"`
				} `json:"event-attendees"`
			} `json:"_links"`
		} `json:"event-organizers"`
		EventAttendees []struct {
			RoleID           string `json:"roleId"`
			RoleName         string `json:"roleName"`
			RoleNamePlural   string `json:"roleNamePlural"`
			RoleDisplayOrder int    `json:"roleDisplayOrder"`
			Links            struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Event struct {
					Href string `json:"href"`
				} `json:"event"`
				Person struct {
					Href string `json:"href"`
				} `json:"person"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"event-attendees"`
		Persons []struct {
			LastName   string `json:"lastName"`
			FirstName  string `json:"firstName"`
			MiddleName string `json:"middleName"`
			FullName   string `json:"fullName"`
			Links      struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"persons"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

type SearchPersonResponse struct {
	Embedded struct {
		Persons []struct {
			LastName   string `json:"lastName"`
			FirstName  string `json:"firstName"`
			MiddleName string `json:"middleName"`
			FullName   string `json:"fullName"`
			Links      struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ID string `json:"id"`
		} `json:"persons"`
		Students []struct {
			ID                string    `json:"id"`
			PersonID          string    `json:"personId"`
			FlowID            string    `json:"flowId"`
			FlowCode          string    `json:"flowCode"`
			SpecialtyCode     string    `json:"specialtyCode"`
			SpecialtyName     string    `json:"specialtyName"`
			SpecialtyProfile  string    `json:"specialtyProfile"`
			LearningStartDate time.Time `json:"learningStartDate"`
			LearningEndDate   time.Time `json:"learningEndDate"`
		} `json:"students"`
	} `json:"_embedded"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}
