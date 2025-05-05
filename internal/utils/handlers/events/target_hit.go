package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
)

type TargetHitHandler struct{}

func (h *TargetHitHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	if len(event.Params) < 1 {
		return fmt.Errorf("missing target parameter for event %d", event.EventID)
	}

	comp.IncrementHits()

	target := event.Params[0]

	logger.LogEvent(event, fmt.Sprintf(entity.MsgTargetHit, target, event.CompetitorID))
	return nil
}
