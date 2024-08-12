package main

import (
	"log"
	"net/http"
	"os"

	"asciiart/web"
)

func main() {
	if len(os.Args) > 1 {
		log.Fatalf("Usage: %s [no arguments required]", os.Args[0])
	}

	mux := http.NewServeMux()

	// Handle static file requests
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Set up route handlers
	mux.HandleFunc("/", web.HomeHandler)
	mux.HandleFunc("/ascii-art", web.AsciiArtHandler)

	log.Println("Starting the server on port 8080")

	// Create an HTTP server with the specified address and handler
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
