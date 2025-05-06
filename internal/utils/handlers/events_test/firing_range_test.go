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

func TestFiringRangeHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &events.FiringRangeHandler{}

	//Test case 1: Competitor enters firing range with valid parameters
	t.Run("valid parameters", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		firingRange := "1"
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventOnFiringRange,
			CompetitorID: competitorID,
			Params:       []string{firingRange},
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventOnFiringRange)
		mockCompetitor.EXPECT().AddFiringRangeStart(timestamp)
		mockCompetitor.EXPECT().SetShotsOnTheFiringRange()
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) is on the firing range(1)"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		//verify
		assert.NoError(t, err)
	})

	// Test case 2: Missing firing range parameter
	t.Run("missing parameter", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 2
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventOnFiringRange,
			CompetitorID: competitorID,
			Params:       []string{},
		}

		err := handler.Handle(event, mockCompetitor, mockLogger)

		//verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing firing range parameter")
	})
}
