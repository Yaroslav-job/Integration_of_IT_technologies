package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"gesture-recognition-results/internal/db"
	"gesture-recognition-results/internal/model"
	"github.com/google/uuid"
)

// HandleRecognition processes POST /recognitions requests.
// Accepts JSON payload, validates and stores it, returning the full result with ID and timestamp.
func HandleRecognition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var result model.RecognitionResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Confidence should be in range 0–1
	if result.Confidence < 0 || result.Confidence > 1 {
		http.Error(w, "Confidence must be 0.0–1.0", http.StatusBadRequest)
		return
	}

	result.ID = uuid.NewString()
	result.Timestamp = time.Now().UTC()

	db.Save(result)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// GetRecognitionsByUser handles GET /recognitions/{userId}.
// Looks up recognition results for the given user.
func GetRecognitionsByUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract user ID from URL
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "User ID required", http.StatusBadRequest)
		return
	}
	userID := parts[2]

	results := db.FindByUserID(userID)
	if results == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(results)
}
