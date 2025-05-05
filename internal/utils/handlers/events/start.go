package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type StartedHandler struct{}

func (h *StartedHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetEventOccurred(entity.EventStarted)
	comp.SetActualStart(&event.Timestamp)

	if comp.GetCurrentLap() == 0 {
		comp.SetCurrentLap(0)
		comp.AddLapStartTime(event.Timestamp)
	}

	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorStarted, event.CompetitorID))
	return nil
}
