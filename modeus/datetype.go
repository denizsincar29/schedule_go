package modeus

import (
	"encoding/json"
	"strings"
	"time"
)

type DateTime time.Time

const (
	isoFullDateTimeTZFormat    = time.RFC3339
	isoFullDateTimeLocalFormat = "2006-01-02T15:04:05" // ISO 8601 without timezone
	isoDateFormat              = "2006-01-02"
)

// MarshalJSON implements the Marshaler interface for DateTime.
func (dt DateTime) MarshalJSON() ([]byte, error) {
	if dt.IsNull() {
		return json.Marshal(nil)
	}
	t := time.Time(dt)
	return json.Marshal(t.Format(isoFullDateTimeTZFormat))
}

// UnmarshalJSON implements the Unmarshaler interface for DateTime.
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	str := string(data)
	// if null, set to zero value
	if str == "null" {
		*dt = DateTime(time.Time{})
		return nil
	}
	// Remove quotes if they exist
	str = strings.Trim(str, "\"")

	var t time.Time
	var err error

	// 1. Try parsing as full ISO 8601 with timezone
	t, err = time.Parse(isoFullDateTimeTZFormat, str)
	if err == nil {
		*dt = DateTime(t)
		return nil
	}

	// 2. Try parsing as full ISO 8601 without timezone (using local timezone)
	t, err = time.Parse(isoFullDateTimeLocalFormat, str)
	if err == nil {
		*dt = DateTime(t)
		return nil
	}

	// 3. Try parsing as ISO 8601 date
	t, err = time.Parse(isoDateFormat, str)
	if err == nil {
		*dt = DateTime(t)
		return nil
	}

	// If all parsing attempts fail, return the last error encountered
	return err
}

// IsNull returns true if the DateTime is the zero value.
func (dt DateTime) IsNull() bool {
	return time.Time(dt).IsZero()
}
