package output

type Writer interface {
	WriteOutputLog(outputLog []string) error
	WriteResults(results []string) error
}
