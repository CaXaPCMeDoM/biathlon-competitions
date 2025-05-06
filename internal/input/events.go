package input

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/infrastructure"
	"biathlon-competitions/internal/parser"
	"fmt"
)

type Event struct {
	reader infrastructure.ReaderEvents
}

func NewEvent(reader infrastructure.ReaderEvents) *Event {
	return &Event{
		reader: reader,
	}
}

func (e *Event) LoadEvents() ([]*entity.Event, error) {
	lines, err := e.reader.ReadLines()
	if err != nil {
		return nil, err
	}

	var events []*entity.Event

	for _, line := range lines {
		event, err := parser.ParseInput(line)
		if err != nil {
			return nil, fmt.Errorf("parse error in line %q: %w", line, err)
		}
		if event != nil {
			events = append(events, event)
		}
	}

	return events, nil
}
