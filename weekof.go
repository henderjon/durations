package durations

import (
	"time"
)

// FmtYearWeek is the format string used by fmt.Println to create a 6 digit string of the yearweek ('YYYYWW').
// e.g. `y, w := t.ISOWeek(); fmt.Sprintf(FmtYearWeek, y, w)`
const FmtYearWeek = "%d%02d"

// Weekof returns the time.Time anchored to the Monday of the week of the given time.Time.
// This is useful for communicating ISO 8601 weeks. See https://golang.org/pkg/time/#Time.ISOWeek
func Weekof(date time.Time) time.Time {

	for int(date.Weekday()) > int(time.Monday) {
		date = date.Add(-time.Hour)
	}

	for int(date.Weekday()) < int(time.Monday) {
		date = date.Add(time.Hour)
	}

	return date
}
