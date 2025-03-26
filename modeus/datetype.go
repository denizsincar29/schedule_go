package modeus

import (
	"log"
	"time"
)

type DateTime time.Time

// marshal unmarshal
func (dt *DateTime) MarshalJSON() ([]byte, error) {
	return time.Time(*dt).MarshalJSON()
}

func (dt *DateTime) UnmarshalJSON(data []byte) error {
	// try to unmarshal as time.Time first
	var t time.Time
	err := t.UnmarshalJSON(data)
	if err == nil {
		*dt = DateTime(t)
		return nil
	}
	// now try to parse without timezone
	t, err = time.Parse("2006-01-02T15:04:05", string(data))
	if err == nil {
		*dt = DateTime(t)
		return nil
	}
	// now try to parse without time completely
	t, err = time.Parse("2006-01-02", string(data))
	if err == nil {
		*dt = DateTime(t)
		return nil
	}
	// don't want any errors here, return zero time
	log.Printf("DateTime.UnmarshalJSON: %v", err)
	*dt = DateTime(time.Time{})
	return nil
}
