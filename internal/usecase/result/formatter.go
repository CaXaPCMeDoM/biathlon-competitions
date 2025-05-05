package result

import (
	"biathlon-competitions/internal/entity"
	"fmt"
	"strings"
	"time"
)

type Formatter struct {
	statusFormatter       func(result *entity.CompetitorResult, sb *strings.Builder)
	idFormatter           func(result *entity.CompetitorResult, sb *strings.Builder)
	lapFormatter          func(result *entity.CompetitorResult, sb *strings.Builder)
	penaltyFormatter      func(result *entity.CompetitorResult, sb *strings.Builder)
	hitsAndShotsFormatter func(result *entity.CompetitorResult, sb *strings.Builder)
}

type FormatterBuilder struct {
	statusFormatter       func(result *entity.CompetitorResult, sb *strings.Builder)
	idFormatter           func(result *entity.CompetitorResult, sb *strings.Builder)
	lapFormatter          func(result *entity.CompetitorResult, sb *strings.Builder)
	penaltyFormatter      func(result *entity.CompetitorResult, sb *strings.Builder)
	hitsAndShotsFormatter func(result *entity.CompetitorResult, sb *strings.Builder)
}

func NewFormatterBuilder() *FormatterBuilder {
	return &FormatterBuilder{
		statusFormatter:       defaultStatusFormatter,
		idFormatter:           defaultIDFormatter,
		lapFormatter:          defaultLapFormatter,
		penaltyFormatter:      defaultPenaltyFormatter,
		hitsAndShotsFormatter: defaultHitsAndShotsFormatter,
	}
}

func (b *FormatterBuilder) WithStatusFormatter(formatter func(result *entity.CompetitorResult, sb *strings.Builder)) *FormatterBuilder {
	b.statusFormatter = formatter
	return b
}

func (b *FormatterBuilder) WithIDFormatter(formatter func(result *entity.CompetitorResult, sb *strings.Builder)) *FormatterBuilder {
	b.idFormatter = formatter
	return b
}

func (b *FormatterBuilder) WithLapFormatter(formatter func(result *entity.CompetitorResult, sb *strings.Builder)) *FormatterBuilder {
	b.lapFormatter = formatter
	return b
}

func (b *FormatterBuilder) WithPenaltyFormatter(formatter func(result *entity.CompetitorResult, sb *strings.Builder)) *FormatterBuilder {
	b.penaltyFormatter = formatter
	return b
}

func (b *FormatterBuilder) WithHitsAndShotsFormatter(formatter func(result *entity.CompetitorResult, sb *strings.Builder)) *FormatterBuilder {
	b.hitsAndShotsFormatter = formatter
	return b
}

func (b *FormatterBuilder) Build() *Formatter {
	return &Formatter{
		statusFormatter:       b.statusFormatter,
		idFormatter:           b.idFormatter,
		lapFormatter:          b.lapFormatter,
		penaltyFormatter:      b.penaltyFormatter,
		hitsAndShotsFormatter: b.hitsAndShotsFormatter,
	}
}

func NewFormatter() *Formatter {
	return NewFormatterBuilder().Build()
}

func (f *Formatter) FormatResults(results []*entity.CompetitorResult) []string {
	var formattedResults []string

	for _, result := range results {
		formattedResults = append(formattedResults, f.formatResult(result))
	}

	return formattedResults
}

func (f *Formatter) formatResult(result *entity.CompetitorResult) string {
	var sb strings.Builder

	f.statusFormatter(result, &sb)
	f.idFormatter(result, &sb)
	f.lapFormatter(result, &sb)
	f.penaltyFormatter(result, &sb)
	f.hitsAndShotsFormatter(result, &sb)

	return sb.String()
}

func defaultStatusFormatter(result *entity.CompetitorResult, sb *strings.Builder) {
	if result.NotStarted {
		sb.WriteString("[NotStarted]")
	} else if result.NotFinished {
		sb.WriteString("[NotFinished]")
	} else {
		sb.WriteString(fmt.Sprintf("[%s]", formatDuration(result.TotalTime)))
	}
}

func defaultIDFormatter(result *entity.CompetitorResult, sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf(" %d", result.ID))
}

func defaultLapFormatter(result *entity.CompetitorResult, sb *strings.Builder) {
	sb.WriteString(" [")
	for i, lapTime := range result.LapTimes {
		if i > 0 {
			sb.WriteString(", ")
		}
		if lapTime > 0 {
			sb.WriteString(fmt.Sprintf("{%s, %.3f}", formatDuration(lapTime), result.LapSpeeds[i]))
		} else {
			sb.WriteString("{,}")
		}
	}
	for i := len(result.LapTimes); i < result.TotalLaps; i++ {
		if i > 0 || len(result.LapTimes) > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("{,}")
	}
	sb.WriteString("]")
}

func defaultPenaltyFormatter(result *entity.CompetitorResult, sb *strings.Builder) {
	if result.PenaltyTime > 0 {
		sb.WriteString(fmt.Sprintf(" {%s, %.3f}", formatDuration(result.PenaltyTime), result.PenaltySpeed))
	} else {
		sb.WriteString(" {,}")
	}
}

func defaultHitsAndShotsFormatter(result *entity.CompetitorResult, sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf(" %d/%d", result.Hits, result.Shots))
}

func formatDuration(d time.Duration) string {
	totalSeconds := int(d.Seconds())
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60
	milliseconds := int(d.Milliseconds()) % 1000

	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)
}
