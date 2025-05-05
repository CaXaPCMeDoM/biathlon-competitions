package file

import (
	"fmt"
	"os"
)

type Writer struct {
	outputLogPath string
	resultsPath   string
}

func NewWriter(outputLogPath, resultsPath string) *Writer {
	return &Writer{
		outputLogPath: outputLogPath,
		resultsPath:   resultsPath,
	}
}

func (w *Writer) WriteOutputLog(outputLog []string) error {
	if w.outputLogPath == "" {
		return nil
	}
	return writeLinesToFile(w.outputLogPath, outputLog)
}

func (w *Writer) WriteResults(results []string) error {
	if w.resultsPath == "" {
		return nil
	}
	return writeLinesToFile(w.resultsPath, results)
}

func writeLinesToFile(path string, lines []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		if _, err := fmt.Fprintln(file, line); err != nil {
			return err
		}
	}

	return nil
}
