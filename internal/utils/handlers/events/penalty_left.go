package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type LeftPenaltyHandler struct{}

func (h *LeftPenaltyHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetEventOccurred(entity.EventLeftPenalty)

	periods := comp.GetPenaltyPeriods()
	if len(periods) == 0 {
		return fmt.Errorf("competitor %d left penalty without entering", event.CompetitorID)
	}

	lastPeriod := periods[len(periods)-1]
	if lastPeriod.EndTime == nil {
		lastPeriod.EndTime = &event.Timestamp
		comp.UpdatePenaltyPeriod(len(periods)-1, lastPeriod)
	}

	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorLeftPenalty, event.CompetitorID))
	return nil
}
