package input

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/parser"
	"biathlon-competitions/internal/port"
	"fmt"
)

type Event struct {
	reader port.ReaderEvents
}

func NewEvent(reader port.ReaderEvents) *Event {
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
