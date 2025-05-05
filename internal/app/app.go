package app

import (
	"biathlon-competitions/internal/adapter/config/json"
	"biathlon-competitions/internal/adapter/events/file"
	"biathlon-competitions/internal/command"
	"biathlon-competitions/internal/output"
	"fmt"
)

func Run(configPath, eventsPath string, writer output.Writer) error {
	configReader := json.New(configPath)
	eventsReader := file.NewReader(eventsPath)

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
