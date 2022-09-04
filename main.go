// the only reasonable approach is to compare created_at dates for each record
// the ones which are newer than a threshold are to be extracted

package main

import (
	"fmt"
	"os"
)

//
func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Args missing. Exiting.")
		os.Exit(1)
	}

	extract_keys(args)
}
