package command

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/infrastructure"
	"biathlon-competitions/internal/input"
	"biathlon-competitions/internal/usecase"
	"biathlon-competitions/internal/usecase/event"
	"biathlon-competitions/internal/usecase/result"
	"fmt"
)

type Handler struct {
	configInput     *input.Config
	eventsInput     *input.Event
	eventProcessor  usecase.EventProcessor
	resultFormatter usecase.ResultFormatter
}

func NewHandler(
	configReader infrastructure.ReaderConfig,
	eventsReader infrastructure.ReaderEvents,
) *Handler {
	return &Handler{
		configInput:     input.NewConfig(configReader),
		eventsInput:     input.NewEvent(eventsReader),
		eventProcessor:  event.NewProcessor(),
		resultFormatter: result.NewFormatter(),
	}
}

func (h *Handler) ProcessEvents() ([]string, []string, error) {
	config, err := h.configInput.LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	events, err := h.eventsInput.LoadEvents()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load events: %w", err)
	}

	results, err := h.eventProcessor.ProcessEvents(events, config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to process events: %w", err)
	}

	outputLog := h.eventProcessor.GenerateOutputLog()

	formattedResults := h.resultFormatter.FormatResults(results)

	return outputLog, formattedResults, nil
}

func (h *Handler) GetCompetitorResults(results []*entity.CompetitorResult) []string {
	return h.resultFormatter.FormatResults(results)
}
