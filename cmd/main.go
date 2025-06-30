package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	r := router.SetupRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error while starting the srever: %v", err)
	}
}
