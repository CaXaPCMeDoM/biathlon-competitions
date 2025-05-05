package entity

const (
	TimeFormatWithMills = "15:04:05.000"
	TimeFormat          = "15:04:05"

	NumberShotsAtTheFiringLine = 5

	EventRegistered      = 1
	EventStartTimeSet    = 2
	EventOnStartLine     = 3
	EventStarted         = 4
	EventOnFiringRange   = 5
	EventTargetHit       = 6
	EventLeftFiringRange = 7
	EventEnteredPenalty  = 8
	EventLeftPenalty     = 9
	EventEndedLap        = 10
	EventCantContinue    = 11

	EventDisqualified = 100
	EventFinished     = 101
)
