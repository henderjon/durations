package durations

import (
	"time"
)

// Nanoseconds returns a duration based on ints. time.Duration is an int64, but there are
// times when the return value of funcs or other operations aren't int64. This is
// pure sugar for casting an int to a time.Duration before doing the math.
func Nanoseconds(n int) time.Duration {
	// 1
	return time.Duration(n) * time.Nanosecond
}

// Microseconds returns a duration based on ints. time.Duration is an int64, but there are
// times when the return value of funcs or other operations aren't int64. This is
// pure sugar for casting an int to a time.Duration before doing the math.
func Microseconds(n int) time.Duration {
	// 1000 * Nanosecond
	return time.Duration(n) * time.Microsecond
}

// Milliseconds returns a duration based on ints. time.Duration is an int64, but there are
// times when the return value of funcs or other operations aren't int64. This is
// pure sugar for casting an int to a time.Duration before doing the math.
func Milliseconds(n int) time.Duration {
	// 1000 * Microsecond
	return time.Duration(n) * time.Millisecond
}

// Seconds returns a duration based on ints. time.Duration is an int64, but there are
// times when the return value of funcs or other operations aren't int64. This is
// pure sugar for casting an int to a time.Duration before doing the math.
func Seconds(n int) time.Duration {
	// 1000 * Millisecond
	return time.Duration(n) * time.Second
}

// Minutes returns a duration based on ints. time.Duration is an int64, but there are
// times when the return value of funcs or other operations aren't int64. This is
// pure sugar for casting an int to a time.Duration before doing the math.
func Minutes(n int) time.Duration {
	// 60 * Second
	return time.Duration(n) * time.Minute
}

// Hours returns a duration based on ints. time.Duration is an int64, but there are
// times when the return value of funcs or other operations aren't int64. This is
// pure sugar for casting an int to a time.Duration before doing the math.
func Hours(n int) time.Duration {
	// 60 * Minute
	return time.Duration(n) * time.Hour
}

// Weekof returns a time.Time of the Monday of the week of the given time.Time.
// This is useful for communicating ISO 8601 weeks.
// read more at https://golang.org/pkg/time/#Time.ISOWeek
func Weekof(date time.Time) time.Time {

	for int(date.Weekday()) > int(time.Monday) {
		date = date.Add(-time.Hour)
	}

	for int(date.Weekday()) < int(time.Monday) {
		date = date.Add(time.Hour)
	}

	return date
}
