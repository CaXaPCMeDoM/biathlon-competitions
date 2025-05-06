package events

import (
	"biathlon-competitions/internal/entity"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCantContinueHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &CantContinueHandler{}

	//Test case 1: Competitor can't continue with a reason
	t.Run("with reason", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		reason := "Lost in the forest"
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventCantContinue,
			CompetitorID: competitorID,
			Params:       []string{reason},
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventCantContinue)
		mockCompetitor.EXPECT().SetNotFinished(true)
		mockCompetitor.EXPECT().SetCantContinue(reason)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) can`t continue: Lost in the forest"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		//verify
		assert.NoError(t, err)
	})

	// Test case 2: Competitor can't continue without a reason
	t.Run("without reason", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 2
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventCantContinue,
			CompetitorID: competitorID,
			Params:       []string{},
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventCantContinue)
		mockCompetitor.EXPECT().SetNotFinished(true)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(2) can`t continue: "))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})
}
