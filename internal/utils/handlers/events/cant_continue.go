package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
	"strings"
)

type CantContinueHandler struct{}

func (h *CantContinueHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	comp.SetEventOccurred(entity.EventCantContinue)

	comp.SetNotFinished(true)
	comment := ""
	if len(event.Params) > 0 {
		comment = strings.Join(event.Params, " ")
		comp.SetCantContinue(comment)
	}
	logger.LogEvent(event, fmt.Sprintf(entity.MsgCompetitorCantContinue, event.CompetitorID, comment))
	return nil
}
