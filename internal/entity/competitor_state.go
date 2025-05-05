package entity

import "time"

type CompetitorState struct {
	ID                int
	PlannedStart      *time.Time
	ActualStart       *time.Time
	LapStartTimes     []time.Time
	LapEndTimes       []time.Time
	PenaltyPeriods    []PenaltyPeriod
	FiringRangeStarts []time.Time
	FiringRangeEnds   []time.Time
	CurrentLap        int
	Hits              int
	Shots             int
	NotStarted        bool
	NotFinished       bool
	Finished          bool
	CantContinue      string
	Registered        bool
	EventState        map[int]bool
}

func NewCompetitorState(id int) *CompetitorState {
	return &CompetitorState{
		ID:         id,
		EventState: make(map[int]bool),
	}
}

func (c *CompetitorState) GetID() int {
	return c.ID
}

func (c *CompetitorState) SetPlannedStart(time *time.Time) {
	c.PlannedStart = time
}

func (c *CompetitorState) SetActualStart(time *time.Time) {
	c.ActualStart = time
}

func (c *CompetitorState) AddLapStartTime(time time.Time) {
	c.LapStartTimes = append(c.LapStartTimes, time)
}

func (c *CompetitorState) AddLapEndTime(time time.Time) {
	c.LapEndTimes = append(c.LapEndTimes, time)
}

func (c *CompetitorState) AddFiringRangeStart(time time.Time) {
	c.FiringRangeStarts = append(c.FiringRangeStarts, time)
}

func (c *CompetitorState) AddFiringRangeEnd(time time.Time) {
	c.FiringRangeEnds = append(c.FiringRangeEnds, time)
}

func (c *CompetitorState) AddPenaltyPeriod(period PenaltyPeriod) {
	c.PenaltyPeriods = append(c.PenaltyPeriods, period)
}

func (c *CompetitorState) GetPenaltyPeriods() []PenaltyPeriod {
	return c.PenaltyPeriods
}

func (c *CompetitorState) UpdatePenaltyPeriod(index int, period PenaltyPeriod) {
	if index >= 0 && index < len(c.PenaltyPeriods) {
		c.PenaltyPeriods[index] = period
	}
}

func (c *CompetitorState) IncrementHits() {
	c.Hits++
}

func (c *CompetitorState) SetShotsOnTheFiringRange() {
	c.Shots += NumberShotsAtTheFiringLine
}

func (c *CompetitorState) SetNotFinished(notFinished bool) {
	c.NotFinished = notFinished
}

func (c *CompetitorState) SetCantContinue(reason string) {
	c.CantContinue = reason
}

func (c *CompetitorState) SetCurrentLap(lap int) {
	c.CurrentLap = lap
}

func (c *CompetitorState) IncrementCurrentLap() {
	c.CurrentLap++
}

func (c *CompetitorState) GetCurrentLap() int {
	return c.CurrentLap
}

func (c *CompetitorState) SetFinished(finished bool) {
	c.Finished = finished
}

func (c *CompetitorState) GetLapEndTimes() []time.Time {
	return c.LapEndTimes
}

func (c *CompetitorState) SetRegistered(registered bool) {
	c.Registered = registered
}

func (c *CompetitorState) IsRegistered() bool {
	return c.Registered
}

func (c *CompetitorState) SetEventOccurred(eventID int) {
	c.EventState[eventID] = true
}

func (c *CompetitorState) HasEventOccurred(eventID int) bool {
	return c.EventState[eventID]
}
