package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type StartLineHandler struct{}

func (h *StartLineHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetEventOccurred(entity.EventOnStartLine)

	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorOnStartLine, event.CompetitorID))
	return nil
}
