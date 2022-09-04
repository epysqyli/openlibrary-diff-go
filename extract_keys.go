package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// args[2] is the input_file, args[3] is the output_file
func extract_keys(args []string) {
	input_file, err := os.Open(args[0])
	log_error(err)
	output_file, err := os.Create(args[1])
	log_error(err)

	parser := csv.NewReader(input_file)
	parser.LazyQuotes = true
	parser.FieldsPerRecord = -1
	parser.Comma = '\t'

	defer output_file.Close()
	defer input_file.Close()

	writer := csv.NewWriter(output_file)
	defer writer.Flush()

	count := 0

	for count < 1 {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}

		log_error(err)
		var entry Entry
		json.Unmarshal([]byte(record[4]), &entry)
		fmt.Println(entry)
		// fmt.Println(entry.Created.Value)
		// writer.Write([]string{record[4]})

		count++
	}
}
