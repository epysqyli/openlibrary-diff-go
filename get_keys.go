package main

import (
	"log"
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
	get_diff("new_keys.txt", "old_keys.txt")
}
