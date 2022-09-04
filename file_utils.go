package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// define the number of concurrent routines
// split the new_keys file into as many partial files as there are routines
// run get_diff() for each partial file, all writing to same dest file

func keys_length() int {
	cmd := exec.Command("wc", "-l", "new_keys.txt")
	result, err := cmd.Output()
	log_error(err)
	new_keys_count, err := strconv.Atoi(strings.Split(string(result), " ")[0])
	log_error(err)
	return new_keys_count
}

func chunk_size(routines_number int, new_keys_length int) int {
	return (keys_length() / routines_number)
}

func split_file(chunks int, chunk_size int) {
	new_keys_file, err := os.Open("new_keys.txt")
	log_error(err)
	new_keys_parser := csv.NewReader(new_keys_file)

	for current := 1; current <= chunks; current++ {
		var base int
		if current == 1 {
			base = 0
		} else {
			base = (current - 1) * chunk_size
		}
		current_limit := chunk_size * current

		partial_diff_keys, err := os.Create(fmt.Sprintf("new_keys_%d.txt", current))
		partial_diff_keys_file := csv.NewWriter(partial_diff_keys)
		log_error(err)

		fmt.Printf("%d => %d %d %s\n", current, base, current_limit, partial_diff_keys.Name())

		for base < current_limit {
			key, err := new_keys_parser.Read()
			log_error(err)
			partial_diff_keys_file.Write(key)
			base++
		}

		partial_diff_keys_file.Flush()
		partial_diff_keys.Close()
	}

	new_keys_file.Close()
}
