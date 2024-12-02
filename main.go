package main

import (
	"log"
	"net/http"
	"os"

	"simple-image-gallery/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/generate", handlers.GenerateImage).Methods("POST", "OPTIONS")
	router.HandleFunc("/images/{id}", handlers.GetImage).Methods("GET", "OPTIONS")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting os port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
