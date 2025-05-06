package events

import (
	"biathlon-competitions/internal/entity"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStartTimeHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &StartTimeHandler{}

	// Test case 1: Valid start time
	t.Run("valid start time", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		startTimeStr := "09:30:00.000"
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventStartTimeSet,
			CompetitorID: competitorID,
			Params:       []string{startTimeStr},
		}

		startTime, _ := time.Parse(entity.TimeFormatWithMills, startTimeStr)

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventStartTimeSet)
		mockCompetitor.EXPECT().SetPlannedStart(gomock.Any()).Do(func(timePtr *time.Time) {
			assert.Equal(t, startTime, *timePtr)
		})
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The start time for the competitor(1) was set by a draw to 09:30:00.000"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})

	// Test case 2: Missing start time parameter
	t.Run("missing start time parameter", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 2
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventStartTimeSet,
			CompetitorID: competitorID,
			Params:       []string{},
		}

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing start time parameter")
	})

	//Test case 3: Invalid start time format
	t.Run("invalid start time format", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 3
		invalidStartTime := "invalid-time"
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventStartTimeSet,
			CompetitorID: competitorID,
			Params:       []string{invalidStartTime},
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventStartTimeSet)

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid start time format")
	})
}
