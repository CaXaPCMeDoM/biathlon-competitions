package events

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"fmt"
	"time"
)

type StartTimeHandler struct{}

func (h *StartTimeHandler) Handle(event *entity.Event, comp handlers.Competitor, logger handlers.EventLogger) error {
	if len(event.Params) < 1 {
		return fmt.Errorf("missing start time parameter for event %d", event.EventID)
	}

	comp.SetEventOccurred(entity.EventStartTimeSet)

	startTimeStr := event.Params[0]
	startTime, err := time.Parse(entity.TimeFormatWithMills, startTimeStr)
	if err != nil {
		return fmt.Errorf("invalid start time format: %v", err)
	}
	comp.SetPlannedStart(&startTime)

	logger.LogEvent(event, fmt.Sprintf(entity.MsgStartTimeSet, event.CompetitorID, startTimeStr))
	return nil
}
