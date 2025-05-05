package entity

import "time"

type CompetitorResult struct {
	ID           int
	NotStarted   bool
	NotFinished  bool
	TotalTime    time.Duration
	LapTimes     []time.Duration
	LapSpeeds    []float64
	PenaltyTime  time.Duration
	PenaltySpeed float64
	Hits         int
	Shots        int
	ActualStart  *time.Time
	PlannedStart *time.Time
	TotalLaps    int
}
