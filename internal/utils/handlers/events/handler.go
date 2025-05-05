package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
)

func GetHandlers() map[int]handlers.EventHandler {
	return map[int]handlers.EventHandler{
		entity.EventRegistered:      &RegistrationHandler{},
		entity.EventStartTimeSet:    &StartTimeHandler{},
		entity.EventOnStartLine:     &StartLineHandler{},
		entity.EventStarted:         &StartedHandler{},
		entity.EventOnFiringRange:   &FiringRangeHandler{},
		entity.EventTargetHit:       &TargetHitHandler{},
		entity.EventLeftFiringRange: &LeftFiringRangeHandler{},
		entity.EventEnteredPenalty:  &EnteredPenaltyHandler{},
		entity.EventLeftPenalty:     &LeftPenaltyHandler{},
		entity.EventEndedLap:        &EndedLapHandler{},
		entity.EventCantContinue:    &CantContinueHandler{},
	}
}
