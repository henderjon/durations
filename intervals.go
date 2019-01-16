package main

import "time"

type Runner func()

func NewRunner(interval time.Duration, runners []Runner) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for _, runner := range runners {
				runner()
			}
		}
	}
}

func Nanoseconds(n int) time.Duration {
	// 1
	return time.Duration(n) * time.Nanosecond
}

func Microseconds(n int) time.Duration {
	// 1000 * Nanosecond
	return time.Duration(n) * time.Microsecond
}

func Milliseconds(n int) time.Duration {
	// 1000 * Microsecond
	return time.Duration(n) * time.Millisecond
}

func Seconds(n int) time.Duration {
	// 1000 * Millisecond
	return time.Duration(n) * time.Second
}

func Minutes(n int) time.Duration {
	// 60 * Second
	return time.Duration(n) * time.Minute
}

func Hours(n int) time.Duration {
	// 60 * Minute
	return time.Duration(n) * time.Hour
}
