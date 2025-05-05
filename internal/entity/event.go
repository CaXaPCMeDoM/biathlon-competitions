package entity

import "time"

type Event struct {
	Timestamp    time.Time
	EventID      int
	CompetitorID int
	Params       []string
}
