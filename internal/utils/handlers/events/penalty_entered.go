package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type EnteredPenaltyHandler struct{}

func (h *EnteredPenaltyHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetEventOccurred(entity.EventEnteredPenalty)

	comp.AddPenaltyPeriod(entity.PenaltyPeriod{
		StartTime: event.Timestamp,
		EndTime:   nil,
	})
	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorEnteredPenalty, event.CompetitorID))
	return nil
}
