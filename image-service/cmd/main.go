package main

import (
	"log"
	"time"

	"monorepo-template/pkg/units"
)

func main() {
	log.Println("Image service started")
	log.Printf("Processing image %d bytes\n", 2*units.MiB)

	<-time.After(time.Minute)

	log.Println("Processing finished, exiting...")
}
