package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type FiringRangeHandler struct{}

func (h *FiringRangeHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	if len(event.Params) < 1 {
		return fmt.Errorf("missing firing range parameter for event %d", event.EventID)
	}

	comp.SetEventOccurred(entity.EventOnFiringRange)

	firingRange := event.Params[0]
	comp.AddFiringRangeStart(event.Timestamp)

	comp.SetShotsOnTheFiringRange()
	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorOnFiringRange, event.CompetitorID, firingRange))
	return nil
}
