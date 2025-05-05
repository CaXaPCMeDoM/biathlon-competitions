package parser

import (
	"biathlon-competitions/internal/entity"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseInput(line string) (*entity.Event, error) {
	line = strings.TrimSpace(line)

	parts := strings.Split(line, "]")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid line: %s", line)
	}

	timeSrt := strings.TrimPrefix(parts[0], "[")

	timestamp, err := time.Parse(entity.TimeFormatWithMills, timeSrt)

	if err != nil {
		return nil, fmt.Errorf("invalid timestamp: %v", err)
	}

	fields := strings.Fields(parts[1])
	if len(fields) < 1 {
		return nil, fmt.Errorf("not enough fields: %s", line)
	}

	eventID, err := strconv.Atoi(fields[0])
	if err != nil {
		return nil, fmt.Errorf("cant parse event id: %v", err)
	}
	competitorID, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, fmt.Errorf("cant parse competitor id: %v", err)
	}

	var params []string
	if len(fields) > 2 {
		params = fields[2:]
	}

	return &entity.Event{
		Timestamp:    timestamp,
		EventID:      eventID,
		CompetitorID: competitorID,
		Params:       params,
	}, nil
}
