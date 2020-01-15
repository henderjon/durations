package durations

import (
	"time"
)

// GetTime is a function that returns a time.Time; it is used here exclusively to inject a specific time
type GetTime func() time.Time

// GetTimeNow fulfills GetTime and is a function that return the current time
func GetTimeNow() time.Time {
	return time.Now()
}

// GetTimeNowUTC fulfills GetTime and is a function that return the current UTC time
func GetTimeNowUTC() time.Time {
	return time.Now().UTC()
}
