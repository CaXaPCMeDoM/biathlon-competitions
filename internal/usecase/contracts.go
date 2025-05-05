package usecase

import "biathlon-competitions/internal/entity"

type (
	EventProcessor interface {
		ProcessEvents(events []*entity.Event, config *entity.Config) ([]*entity.CompetitorResult, error)
		GenerateOutputLog() []string
	}

	ResultFormatter interface {
		FormatResults(results []*entity.CompetitorResult) []string
	}
)
