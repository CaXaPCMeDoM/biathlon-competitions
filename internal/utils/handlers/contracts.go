package handlers

import (
	"biathlon-competitions/internal/entity"
	"time"
)

//go:generate mockgen -source=contracts.go -destination=mocks/mock.go

type EventHandler interface {
	Handle(event *entity.Event, comp Competitor, logger EventLogger) error
}

type Competitor interface {
	GetID() int
	SetPlannedStart(time *time.Time)
	SetActualStart(time *time.Time)
	AddLapStartTime(time time.Time)
	AddLapEndTime(time time.Time)
	AddFiringRangeStart(time time.Time)
	AddFiringRangeEnd(time time.Time)
	AddPenaltyPeriod(period entity.PenaltyPeriod)
	GetPenaltyPeriods() []entity.PenaltyPeriod
	UpdatePenaltyPeriod(index int, period entity.PenaltyPeriod)
	IncrementHits()
	SetShotsOnTheFiringRange()
	SetNotFinished(notFinished bool)
	SetCantContinue(reason string)
	SetCurrentLap(lap int)
	IncrementCurrentLap()
	GetCurrentLap() int
	SetFinished(finished bool)
	GetLapEndTimes() []time.Time
	SetRegistered(registered bool)
	IsRegistered() bool
	SetEventOccurred(eventID int)
	HasEventOccurred(eventID int) bool
}

type EventLogger interface {
	LogEvent(event *entity.Event, message string)
}

type LapCompletionChecker interface {
	CheckLapCompletion(comp Competitor)
}
