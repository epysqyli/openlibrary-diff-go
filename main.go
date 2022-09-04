package main

import (
	"fmt"
	"os"
)

// example date: "2022-03-01T00:00:00Z00:00"
func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Printf("3 parameters needed:\nInput file of records\nOutput file name\nThreshold datetime\n")
		os.Exit(1)
	}

	extract_new_records(args)
}
