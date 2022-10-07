package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
)

// args[0] input_file, args[1] output_file
func extract_new_records(args []string, threshold string) {
	input_file, err := os.Open(args[0])
	log_error(err)
	output_file, err := os.Create(args[1])
	log_error(err)

	parser := csv.NewReader(input_file)
	parser.LazyQuotes = true
	parser.FieldsPerRecord = 5
	parser.Comma = '\t'

	for {
		record, err := parser.Read()
		log_error(err)
		if err == io.EOF {
			break
		}

		var entry Entry
		json.Unmarshal([]byte(record[len(record)-1]), &entry)
		record_is_new := entry.Created.Value > threshold
		if record_is_new == true {
			output_file.WriteString(record[len(record)-1])
			output_file.WriteString("\n")
		}
	}

	output_file.Sync()
	output_file.Close()
	input_file.Close()
}
