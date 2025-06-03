package main

import (
	"log"
	"net/http"

	"gesture-recognition-results/internal/handler"
)

// main starts the gesture-results HTTP service.
// Registers routes and begins listening on port 8080.
func main() {
	http.HandleFunc("/recognitions", handler.HandleRecognition)
	http.HandleFunc("/recognitions/", handler.GetRecognitionsByUser)

	log.Println("Starting gesture-results service on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
