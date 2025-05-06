package file

import (
	"bufio"
	"os"
)

type EventReader struct {
	filepath string
}

func NewEventReader(filepathEvents string) *EventReader {
	return &EventReader{
		filepath: filepathEvents,
	}
}

func (r *EventReader) ReadLines() ([]string, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
