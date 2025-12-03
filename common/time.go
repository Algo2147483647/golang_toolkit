// Package common provides common utility functions for time handling
package common

import "time"

// MaxTime returns the later of two time.Time values
func MaxTime(t1, t2 time.Time) time.Time {
	if t1.After(t2) {
		return t1
	}
	return t2
}

// MinTime returns the earlier of two time.Time values
func MinTime(t1, t2 time.Time) time.Time {
	if t1.Before(t2) {
		return t1
	}
	return t2
}

// CompareDate compares only the date (year, month, day) of two time.Time values.
// Returns -1 if t1 is before t2, 1 if t1 is after t2, and 0 if they are equal.
func CompareDate(t1, t2 time.Time) int {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	if y1 < y2 {
		return -1
	} else if y1 > y2 {
		return 1
	}

	if m1 < m2 {
		return -1
	} else if m1 > m2 {
		return 1
	}

	if d1 < d2 {
		return -1
	} else if d1 > d2 {
		return 1
	}

	return 0
}

// CombineDateAndClock combines the date part from one time value and
// the clock part (hours, minutes, seconds, nanoseconds) from another.
func CombineDateAndClock(date, clock time.Time) time.Time {
	return time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		clock.Hour(),
		clock.Minute(),
		clock.Second(),
		clock.Nanosecond(),
		date.Location(),
	)
}

func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}
