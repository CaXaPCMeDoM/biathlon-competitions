package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type RegistrationHandler struct{}

func (h *RegistrationHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetRegistered(true)
	comp.SetEventOccurred(entity.EventRegistered)
	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorRegistered, event.CompetitorID))
	return nil
}
