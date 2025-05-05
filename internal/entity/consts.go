package entity

const (
	TimeFormatWithMills = "15:04:05.000"
	TimeFormat          = "15:04:05"

	NumberShotsAtTheFiringLine = 5
)

const (
	EventDisqualified = 100
	EventFinished     = 101
)

const (
	EventRegistered = iota + 1
	EventStartTimeSet
	EventOnStartLine
	EventStarted
	EventOnFiringRange
	EventTargetHit
	EventLeftFiringRange
	EventEnteredPenalty
	EventLeftPenalty
	EventEndedLap
	EventCantContinue
)
