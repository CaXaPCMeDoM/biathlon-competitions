package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type EndedLapHandler struct{}

func (h *EndedLapHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetEventOccurred(entity.EventEndedLap)

	comp.AddLapEndTime(event.Timestamp)
	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorEndedLap, event.CompetitorID))

	if checker, ok := logger.(handlers.LapCompletionChecker); ok {
		checker.CheckLapCompletion(comp)
	}

	return nil
}
