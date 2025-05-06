package events

import (
	"biathlon-competitions/internal/entity"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockLoggerWithChecker struct {
	*mocks.MockEventLogger
	*mocks.MockLapCompletionChecker
}

func TestEndedLapHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handler := &EndedLapHandler{}

	//Test case 1: Logger implements LapCompletionChecker
	t.Run("logger implements LapCompletionChecker", func(t *testing.T) {
		mockCompetitor := mocks.NewMockCompetitor(ctrl)
		mockLogger := mocks.NewMockEventLogger(ctrl)
		mockChecker := mocks.NewMockLapCompletionChecker(ctrl)

		combinedMock := &mockLoggerWithChecker{
			MockEventLogger:          mockLogger,
			MockLapCompletionChecker: mockChecker,
		}

		timestamp := time.Now()
		competitorID := 1
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventEndedLap,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventEndedLap)
		mockCompetitor.EXPECT().AddLapEndTime(timestamp)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) ended the main lap"))
		mockChecker.EXPECT().CheckLapCompletion(mockCompetitor)

		err := handler.Handle(event, mockCompetitor, combinedMock)

		// verify
		assert.NoError(t, err)
	})

	// Test case 2: Logger does not implement LapCompletionChecker
	t.Run("logger does not implement LapCompletionChecker", func(t *testing.T) {
		mockCompetitor := mocks.NewMockCompetitor(ctrl)
		mockLogger := mocks.NewMockEventLogger(ctrl)

		timestamp := time.Now()
		competitorID := 2
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventEndedLap,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventEndedLap)
		mockCompetitor.EXPECT().AddLapEndTime(timestamp)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(2) ended the main lap"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})
}
