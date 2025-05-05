package file

import (
	"bufio"
	"os"
)

type Reader struct {
	filepath string
}

func NewReader(filepathEvents string) *Reader {
	return &Reader{
		filepath: filepathEvents,
	}
}

func (r *Reader) ReadLines() ([]string, error) {
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
