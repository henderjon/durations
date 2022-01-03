package durations

import (
	"time"
)

// Weekof returns the time.Time anchored to the Monday of the week of the given time.Time.
// This is useful for communicating ISO 8601 weeks. See https://golang.org/pkg/time/#Time.ISOWeek
func Weekof(date time.Time) time.Time {

	for date.Weekday() > time.Monday {
		date = date.Add(-time.Hour)
	}

	for date.Weekday() < time.Monday {
		date = date.Add(time.Hour)
	}

	return date
}
