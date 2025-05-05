package main

import (
	"biathlon-competitions/internal/app"
	"biathlon-competitions/internal/output"
	"biathlon-competitions/internal/output/console"
	"biathlon-competitions/internal/output/file"
	"flag"
	"fmt"
	"os"
)

func main() {
	configPath := flag.String("config", "./sunny_5_skiers/config.json", "Path to configuration file")
	eventsPath := flag.String("events", "./sunny_5_skiers/events", "Path to events file")
	outputLogPath := flag.String("output-log", "./result/output.log", "Path to output log file (optional)")
	resultsPath := flag.String("results", "./result/results", "Path to results file (optional)")

	flag.Parse()

	var writer output.Writer
	if *outputLogPath != "" || *resultsPath != "" {
		writer = file.NewWriter(*outputLogPath, *resultsPath)
	} else {
		writer = console.NewWriter()
	}

	err := app.Run(*configPath, *eventsPath, writer)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
