package app

import (
	"biathlon-competitions/internal/command"
	"biathlon-competitions/internal/infrastructure/config/json"
	"biathlon-competitions/internal/infrastructure/events/file"
	"biathlon-competitions/internal/output"
	"fmt"
)

func Run(configPath, eventsPath string, writer output.Writer) error {
	configReader := json.NewConfigReader(configPath)
	eventsReader := file.NewEventReader(eventsPath)

	handler := command.NewHandler(configReader, eventsReader)

	outputLog, results, err := handler.ProcessEvents()
	if err != nil {
		return fmt.Errorf("error processing events: %w", err)
	}

	if err := writeOutput(writer, outputLog, results); err != nil {
		return err
	}

	return nil
}

func writeOutput(writer output.Writer, outputLog, results []string) error {
	if err := writer.WriteOutputLog(outputLog); err != nil {
		return fmt.Errorf("error writing output log: %w", err)
	}

	if err := writer.WriteResults(results); err != nil {
		return fmt.Errorf("error writing results: %w", err)
	}

	return nil
}
