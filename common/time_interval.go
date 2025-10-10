package common

import "time"

// TimeInterval represents a time interval with a start and end time
type TimeInterval struct {
	StartTime time.Time
	EndTime   time.Time
}

// InTimeInterval checks if a given time is within the time interval (inclusive)
func (ti *TimeInterval) InTimeInterval(t time.Time) bool {
	return t.Before(ti.StartTime) || t.After(ti.EndTime)
}

// IsZero checks if the TimeInterval is zero (both start and end times are zero)
func (ti *TimeInterval) IsZero() bool {
	return ti.StartTime.IsZero() && ti.EndTime.IsZero()
}

// Duration returns the duration of the time interval
func (ti *TimeInterval) Duration() time.Duration {
	return ti.EndTime.Sub(ti.StartTime)
}
