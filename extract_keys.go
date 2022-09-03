package main

import (
	"encoding/csv"
	"io"
	"os"
	"strings"
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

	defer output_file.Close()
	defer input_file.Close()

	writer := csv.NewWriter(output_file)
	defer writer.Flush()

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}

		log_error(err)

		entries := strings.Fields(record[0])
		key_slice := strings.Split(entries[1], "/")
		key := []string{key_slice[len(key_slice)-1]}
		writer.Write(key)
	}
}
