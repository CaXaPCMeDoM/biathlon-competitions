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

func TestTargetHitHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCompetitor := mocks.NewMockCompetitor(ctrl)
	mockLogger := mocks.NewMockEventLogger(ctrl)

	handler := &events.TargetHitHandler{}

	// Test case 1: Valid target parameter
	t.Run("valid target parameter", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 1
		target := "1"
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventTargetHit,
			CompetitorID: competitorID,
			Params:       []string{target},
		}

		mockCompetitor.EXPECT().IncrementHits()
		mockLogger.EXPECT().LogEvent(event, gomock.Eq(
			"The target(1) has been hit by competitor(1)"))

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.NoError(t, err)
	})

	// Test case 2: Missing target parameter
	t.Run("missing target parameter", func(t *testing.T) {
		timestamp := time.Now()
		competitorID := 2
		event := &entity.Event{
			Timestamp:    timestamp,
			EventID:      entity.EventTargetHit,
			CompetitorID: competitorID,
			Params:       []string{},
		}

		err := handler.Handle(event, mockCompetitor, mockLogger)

		// verify
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing target parameter")
	})
}
