package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("2 parameters needed:\nInput file of records\nOutput file name\n")
		os.Exit(1)
	}

	if len(args) == 2 {
		extract_new_records(args, time.Now().Format(time.RFC3339))
	}

	// example date: "2022-03-01T00:00:00Z00:00"
	// supply time RFC3339 compliant time threshold from as cmd arg
	if len(args) == 3 {
		extract_new_records(args, args[2])
	}
}
