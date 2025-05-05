package event

import (
	"biathlon-competitions/internal/entity"
	"biathlon-competitions/internal/utils/handlers"
	"biathlon-competitions/internal/utils/handlers/events"
	"fmt"
	"sort"
	"time"
)

type Processor struct {
	events        []*entity.Event
	config        *entity.Config
	competitors   map[int]*entity.CompetitorState
	outputLog     []string
	eventHandlers map[int]handlers.EventHandler
}

func NewProcessor() *Processor {
	return &Processor{
		competitors:   make(map[int]*entity.CompetitorState),
		eventHandlers: events.GetHandlers(),
	}
}

func (p *Processor) ProcessEvents(events []*entity.Event, config *entity.Config) ([]*entity.CompetitorResult, error) {
	p.events = events
	p.config = config
	p.outputLog = []string{}

	for _, event := range events {
		if err := p.processEvent(event); err != nil {
			return nil, err
		}
	}

	p.checkNotStarted()

	return p.generateResults(), nil
}

func (p *Processor) GenerateOutputLog() []string {
	return p.outputLog
}

func (p *Processor) processEvent(event *entity.Event) error {
	competitorID := event.CompetitorID
	eventID := event.EventID

	comp, exists := p.competitors[competitorID]
	if !exists {
		comp = entity.NewCompetitorState(competitorID)
		p.competitors[competitorID] = comp
	}

	eventHandler, exists := p.eventHandlers[eventID]
	if !exists {
		return fmt.Errorf("unknown event ID: %d", eventID)
	}

	if !comp.IsRegistered() && eventID != entity.EventRegistered {
		return nil
	}

	if eventID == entity.EventStarted && !comp.HasEventOccurred(entity.EventOnStartLine) {
		comp.NotStarted = true
		p.LogEvent(&entity.Event{
			Timestamp:    event.Timestamp,
			EventID:      entity.EventDisqualified,
			CompetitorID: competitorID,
		}, fmt.Sprintf(entity.MsgCompetitorDisqualified, competitorID))
		return nil
	}

	return eventHandler.Handle(event, comp, p)
}

func (p *Processor) LogEvent(event *entity.Event, message string) {
	logEntry := fmt.Sprintf("[%s] %s", event.Timestamp.Format(entity.TimeFormatWithMills), message)
	p.outputLog = append(p.outputLog, logEntry)
}

func (p *Processor) checkNotStarted() {
	for _, comp := range p.competitors {
		if !comp.Registered {
			continue
		}

		if comp.PlannedStart != nil && comp.ActualStart == nil {
			comp.NotStarted = true

			p.LogEvent(&entity.Event{
				Timestamp:    *comp.PlannedStart,
				EventID:      entity.EventDisqualified,
				CompetitorID: comp.ID,
			}, fmt.Sprintf(entity.MsgCompetitorDisqualified, comp.ID))
		} else if comp.PlannedStart != nil && comp.ActualStart != nil {
			maxStartTime := comp.PlannedStart.Add(p.config.StartDelta.Duration)
			if comp.ActualStart.After(maxStartTime) {
				comp.NotStarted = true

				p.LogEvent(&entity.Event{
					Timestamp:    *comp.ActualStart,
					EventID:      entity.EventDisqualified,
					CompetitorID: comp.ID,
				}, fmt.Sprintf(entity.MsgCompetitorDisqualified, comp.ID))
			}
		}
	}
}

func (p *Processor) generateResults() []*entity.CompetitorResult {
	var results []*entity.CompetitorResult

	for id, comp := range p.competitors {
		if !comp.Registered {
			continue
		}

		result := &entity.CompetitorResult{
			ID:          id,
			NotStarted:  comp.NotStarted,
			NotFinished: comp.NotFinished,
			Hits:        comp.Hits,
			Shots:       comp.Shots,
			TotalLaps:   p.config.Laps,
		}

		if comp.PlannedStart != nil {
			result.PlannedStart = comp.PlannedStart
		}
		if comp.ActualStart != nil {
			result.ActualStart = comp.ActualStart
		}

		if !comp.NotStarted && len(comp.LapEndTimes) > 0 {
			for i := 0; i < len(comp.LapEndTimes); i++ {
				lapTime := comp.LapEndTimes[i].Sub(comp.LapStartTimes[i])
				result.LapTimes = append(result.LapTimes, lapTime)

				speed := float64(p.config.LapLen) / lapTime.Seconds()
				result.LapSpeeds = append(result.LapSpeeds, speed)
			}

			if comp.ActualStart != nil && comp.PlannedStart != nil {
				startDelay := comp.ActualStart.Sub(*comp.PlannedStart)
				if len(comp.LapEndTimes) == p.config.Laps {
					result.TotalTime = comp.LapEndTimes[len(comp.LapEndTimes)-1].Sub(*comp.ActualStart) + startDelay
				} else {
					result.TotalTime = comp.LapEndTimes[len(comp.LapEndTimes)-1].Sub(*comp.ActualStart) + startDelay
				}
			}
		}

		var totalPenaltyTime time.Duration
		penaltyLaps := 0
		for _, period := range comp.PenaltyPeriods {
			if period.EndTime != nil {
				totalPenaltyTime += period.EndTime.Sub(period.StartTime)
				penaltyLaps++
			}
		}

		result.PenaltyTime = totalPenaltyTime
		if totalPenaltyTime > 0 {
			totalPenaltyLen := penaltyLaps * p.config.PenaltyLen
			result.PenaltySpeed = float64(totalPenaltyLen) / totalPenaltyTime.Seconds()
		} else {
			result.PenaltySpeed = 0
		}

		if comp.CurrentLap < p.config.Laps {
			result.NotFinished = true
		}

		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].NotStarted != results[j].NotStarted {
			return !results[i].NotStarted
		}

		if results[i].NotFinished != results[j].NotFinished {
			return !results[i].NotFinished
		}

		return results[i].TotalTime < results[j].TotalTime
	})

	return results
}

func (p *Processor) CheckLapCompletion(comp handlers.Competitor) {
	comp.IncrementCurrentLap()
	if len(comp.GetLapEndTimes()) == p.config.Laps {
		comp.SetFinished(true)
		p.LogEvent(&entity.Event{
			Timestamp:    comp.GetLapEndTimes()[len(comp.GetLapEndTimes())-1],
			EventID:      entity.EventFinished,
			CompetitorID: comp.GetID(),
		}, fmt.Sprintf(entity.MsgCompetitorFinished, comp.GetID()))
	} else {
		comp.AddLapStartTime(comp.GetLapEndTimes()[len(comp.GetLapEndTimes())-1])
	}
}
