package jobs

import (
	"bufio"
	"os"

	"github.com/goravel/framework/facades"
)

type CsvReaderJob struct {
}

// Signature The name and signature of the job.
func (receiver *CsvReaderJob) Signature() string {
	return "csv_reader_job"
}

// Handle Execute the job.
func (receiver *CsvReaderJob) Handle(args ...any) error {
	// Get filename from args
	filename := args[0].(string)

	// Get path to the file
	path := facades.Storage().Path("csv/" + filename)

	csv, err := os.Open(path)

	// Handle error while opening file
	if err != nil {
		return err
	}

	defer csv.Close()

	var content string
	// Read the file content
	scanner := bufio.NewScanner(csv)

	for scanner.Scan() {
		content += scanner.Text() + "\n"
		println(scanner.Text())
	}

	return nil
}
