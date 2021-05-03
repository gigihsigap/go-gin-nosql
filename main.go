package main

import (
	"log"
	"os"

	"backend/api"
)

func main() {
	port := os.Getenv("PORT")
	server := api.NewServer()

	if port == "" {
		port = "8000"
	}

	err := server.Start(":" + port)

	if err != nil {
		log.Fatal("Cannot start server!", err)
	}
}
