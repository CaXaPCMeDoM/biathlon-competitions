package events

import (
	"biathlon-competitions/internal/entity"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLeftPenaltyHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &LeftPenaltyHandler{}

	//Test case 1: Competitor leaves penalty lap after entering one
	t.Run("competitor leaves penalty lap after entering", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventLeftPenalty,
			CompetitorID: competitorID,
		}

		startTime := timestamp.Add(-time.Minute) // <started 1 minute ago>
		penaltyPeriod := entity.PenaltyPeriod{
			StartTime: startTime,
			EndTime:   nil,
		}
		periods := []entity.PenaltyPeriod{penaltyPeriod}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventLeftPenalty)
		mockCompetitor.EXPECT().GetPenaltyPeriods().Return(periods)
		mockCompetitor.EXPECT().UpdatePenaltyPeriod(0, gomock.Any()).Do(func(index int, period entity.PenaltyPeriod) {
			assert.Equal(t, 0, index)
			assert.Equal(t, startTime, period.StartTime)
			assert.Equal(t, timestamp, *period.EndTime)
		})
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) left the penalty laps"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})

	//Test case 2: Competitor tries to leave penalty lap without entering one
	t.Run("competitor leaves penalty lap without entering", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 2
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventLeftPenalty,
			CompetitorID: competitorID,
		}

		var periods []entity.PenaltyPeriod

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventLeftPenalty)
		mockCompetitor.EXPECT().GetPenaltyPeriods().Return(periods)

		err := handler.Handle(event, mockCompetitor, mockLogger)

		//verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "competitor 2 left penalty without entering")
	})
}
