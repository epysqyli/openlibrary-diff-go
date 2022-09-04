package main

import (
	"encoding/csv"
	"io"
	"os"
)

func get_diff(new_keys string, old_keys string, dest *csv.Writer) {
	new_keys_file, err := os.Open(new_keys)
	log_error(err)
	new_keys_parser := csv.NewReader(new_keys_file)

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
				continue
			}
		}

		if is_match == false {
			dest.Write(new_key)
		}
	}
}
