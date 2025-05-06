package events

import (
	"biathlon-competitions/internal/entity"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRegistrationHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &RegistrationHandler{}

	// Test case: Competitor registers
	t.Run("competitor registers", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventRegistered,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetRegistered(true)
		mockCompetitor.EXPECT().SetEventOccurred(entity.EventRegistered)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) registered"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})
}
