package main

import (
	"log"
	"time"

	"monorepo-template/pkg/units"
)

func main() {
	log.Println("Worker started")
	log.Printf("Processing message %d bytes\n", 2*units.GiB)

	<-time.After(time.Minute)

	log.Println("Processing finished, exiting...")
}
