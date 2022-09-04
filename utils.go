package main

import (
	"log"
)

func log_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
