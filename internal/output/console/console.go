package console

import (
	"fmt"
)

type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) WriteOutputLog(outputLog []string) error {
	fmt.Println("Output log:")
	for _, line := range outputLog {
		fmt.Println(line)
	}
	return nil
}

func (w *Writer) WriteResults(results []string) error {
	fmt.Println("\nResulting table:")
	for _, line := range results {
		fmt.Println(line)
	}
	return nil
}
