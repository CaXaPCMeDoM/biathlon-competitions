package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type LeftFiringRangeHandler struct{}

func (h *LeftFiringRangeHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetEventOccurred(entity.EventLeftFiringRange)

	comp.AddFiringRangeEnd(event.Timestamp)

	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorLeftFiringRange, event.CompetitorID))
	return nil
}
