package main

import (
	"log"
	"net/http"

	"github.com/eugeneradionov/monorepo-template/api/handlers"
)

func main() {
	err := http.ListenAndServe(":8080", handlers.New())
	if err != nil {
		log.Fatalf("http server failed: %v", err)
	}

	log.Println("api exited")
}
