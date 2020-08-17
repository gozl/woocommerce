package models

import (
	"time"
	"errors"
)

// Timestamp represents a timestamp
type Timestamp time.Time

// MarshalJSON serializes a Timestamp into JSON encoded data
func (c *Timestamp) MarshalJSON() ([]byte, error) {
	t := time.Time(*c)

	if t.IsZero() {
		return []byte("null"), nil
	}
	
	tStr := t.UTC().Format(time.RFC3339)
	if tStr[len(tStr)-1] != 'Z' {
		return nil, errors.New("unable to marshal timestamp to UTC")
	}
	tStr = tStr[0:len(tStr)-1]

	return []byte("\"" + tStr + "\""), nil
}

// UnmarshalJSON parses JSON encoded data to a Timestamp
func (c *Timestamp) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	timeStr := string(data)
	if timeStr == "null" {
		return nil
	}

	if len(timeStr) < 3 || timeStr[0] != '"' || timeStr[len(timeStr)-1] != '"' {
		return errors.New("time should be formatted as string in JSON: " + timeStr)
	}
	timeStr = timeStr[1:len(timeStr)-1] + "Z"

	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return err
	}

	*c = Timestamp(t)
	return nil
}