package entity

import "time"

type PenaltyPeriod struct {
	StartTime time.Time
	EndTime   *time.Time
}
