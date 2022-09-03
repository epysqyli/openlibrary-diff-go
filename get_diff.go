package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func get_diff(new_keys string, old_keys string) {
	new_keys_file, err := os.Open(new_keys)
	log_error(err)
	new_keys_parser := csv.NewReader(new_keys_file)

	diff_output, err := os.Create("diff_keys.txt")
	log_error(err)

	diff_writer := csv.NewWriter(diff_output)

	for {
		new_key, err := new_keys_parser.Read()
		if err == io.EOF {
			break
		}

		log_error(err)
		is_match := false

		old_keys_file, err := os.Open(old_keys)
		log_error(err)
		old_keys_parser := csv.NewReader(old_keys_file)

		for {
			old_key, err := old_keys_parser.Read()
			if err == io.EOF {
				break
			}

			log_error(err)

			if new_key[0] == old_key[0] {
				is_match = true
				fmt.Printf("MATCH -> new_key: %s - old_key: %s\n", new_key[0], old_key[0])
				continue
			}
		}

		if is_match == false {
			diff_writer.Write(new_key)
		}
	}

	diff_writer.Flush()
}
