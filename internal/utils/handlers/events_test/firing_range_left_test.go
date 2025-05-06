package events_test

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers/events"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLeftFiringRangeHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &events.LeftFiringRangeHandler{}

	// Test case: Competitor leaves firing range
	t.Run("competitor leaves firing range", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventLeftFiringRange,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventLeftFiringRange)
		mockCompetitor.EXPECT().AddFiringRangeEnd(timestamp)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) left the firing range"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})
}
