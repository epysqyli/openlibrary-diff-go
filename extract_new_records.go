package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
)

// args[0] is the input_file
// args[1] is the output_file
func extract_new_records(args []string, threshold string) {
	input_file, err := os.Open(args[0])
	log_error(err)
	output_file, err := os.Create(args[1])
	log_error(err)

	parser := csv.NewReader(input_file)
	parser.LazyQuotes = true
	parser.FieldsPerRecord = -1
	parser.Comma = '\t'

	writer := csv.NewWriter(output_file)

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}

		log_error(err)
		var entry Entry
		json.Unmarshal([]byte(record[4]), &entry)

		record_is_new := entry.Created.Value > threshold
		if record_is_new == true {
			writer.Write(record)
		}
	}

	writer.Flush()
	output_file.Close()
	input_file.Close()
}
