package events

import (
	"biathlon-competitions/internal/entity"
	mocks "biathlon-competitions/internal/utils/handlers/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStartLineHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &StartLineHandler{}

	// Test case: Competitor is on the start line
	t.Run("competitor is on the start line", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventOnStartLine,
			CompetitorID: competitorID,
		}

		mockCompetitor.EXPECT().SetEventOccurred(entity.EventOnStartLine)
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The competitor(1) is on the start line"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})
}
