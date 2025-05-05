package entity

import (
	"fmt"
	"time"
)

type Config struct {
	Laps        int            `json:"laps"`
	LapLen      int            `json:"lapLen"`
	PenaltyLen  int            `json:"penaltyLen"`
	FiringLines int            `json:"firingLines"`
	Start       customTime     `json:"start"`
	StartDelta  customDuration `json:"startDelta"`
}

type customTime struct {
	time.Time
}

func (ct *customTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]

	t, err := time.Parse(TimeFormatWithMills, s)
	if err != nil {
		t, err = time.Parse(TimeFormat, s)
		if err != nil {
			return err
		}
	}
	ct.Time = t
	return nil
}

type customDuration struct {
	time.Duration
}

func (cd *customDuration) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]

	d, err := time.ParseDuration(parseToDurationFormat(s))
	if err != nil {
		return err
	}
	cd.Duration = d
	return nil
}

func parseToDurationFormat(s string) string {
	t, _ := time.Parse(TimeFormat, s)
	h := t.Hour()
	m := t.Minute()
	sec := t.Second()

	duration := ""
	if h > 0 {
		duration += fmt.Sprintf("%dh", h)
	}
	if m > 0 {
		duration += fmt.Sprintf("%dm", m)
	}
	if sec > 0 || duration == "" {
		duration += fmt.Sprintf("%ds", sec)
	}
	return duration
}
