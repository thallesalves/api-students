package main

import (
	"log"

	"github.com/thallesalves/api-students/api"
)

func main() {

	server := api.NewServer()
	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
