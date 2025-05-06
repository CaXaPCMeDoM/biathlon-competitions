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

func TestStartedHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &events.StartedHandler{}

	// Test case 1: Competitor starts with current lap 0 (first lap)
	t.Run("competitor starts first lap", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventStarted,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventStarted)
		mockCompetitor.EXPECT().SetActualStart(&event.Timestamp)
		mockCompetitor.EXPECT().GetCurrentLap().Return(0)
		mockCompetitor.EXPECT().SetCurrentLap(0)
		mockCompetitor.EXPECT().AddLapStartTime(timestamp)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) has started"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})

	// Test case 2: Competitor starts with current lap > 0 (not first lap)
	t.Run("competitor starts not first lap", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 2
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventStarted,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventStarted)
		mockCompetitor.EXPECT().SetActualStart(&event.Timestamp)
		mockCompetitor.EXPECT().GetCurrentLap().Return(1) // !! already on lap 1 !!!
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(2) has started"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})
}
