package main

import (
	"book-store/cmd/bookstrap"
	"log"
)

func main() {
	router := bookstrap.NewRouter()

	if err := router.Run(":9000"); err != nil {
		log.Fatalf("cloud not start server: %v", err)
	}
}
