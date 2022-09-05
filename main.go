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
		now := time.Now()
		last_month := time.Month(now.Month() - 1)
		last_month_time := time.Date(now.Year(), last_month, 1, 0, 0, 0, 0, time.Local)
		extract_new_records(args, last_month_time.Format(time.RFC3339))
	}

	// example date: "2022-03-01T00:00:00Z00:00"
	// supply time RFC3339 compliant time threshold from as cmd arg
	if len(args) == 3 {
		extract_new_records(args, args[2])
	}
}
