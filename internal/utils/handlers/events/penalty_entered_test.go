package events

import (
	"biathlon-competitions/internal/entity"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestEnteredPenaltyHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &EnteredPenaltyHandler{}

	//Test case: Competitor enters penalty lap
	t.Run("competitor enters penalty lap", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventEnteredPenalty,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventEnteredPenalty)
		mockCompetitor.EXPECT().AddPenaltyPeriod(gomock.Any()).Do(func(period entity.PenaltyPeriod) {
			assert.Equal(t, timestamp, period.StartTime)
			assert.Nil(t, period.EndTime)
		})
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) entered the penalty laps"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})
}
