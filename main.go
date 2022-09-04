package main

import (
	"encoding/csv"
	"log"
	"os"
)

func log_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// args := os.Args[1:]
	// if len(args) < 2 {
	// 	fmt.Println("Args missing. Exiting.")
	// 	os.Exit(1)
	// }
	// extract_keys(args)

	diff_keys_output, err := os.Create("diff_keys.txt")
	log_error(err)
	dest := csv.NewWriter(diff_keys_output)

	// new_keys_length := keys_length()
	// chunks := 6
	// chunk_size := chunk_size(chunks, new_keys_length)
	// split_file(chunks, chunk_size)
	get_diff("new_keys_1.txt", "old_keys.txt", dest)
}
