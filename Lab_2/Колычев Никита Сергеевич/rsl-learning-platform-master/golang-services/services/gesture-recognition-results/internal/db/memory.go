package db

import (
	"gesture-recognition-results/internal/model"
)

// memoryStore is an in-memory map of user ID to gesture recognition results.
var memoryStore = make(map[string][]model.RecognitionResult)

// Save stores a new recognition result for a user.
func Save(result model.RecognitionResult) {
	memoryStore[result.UserID] = append(memoryStore[result.UserID], result)
}

// FindByUserID retrieves all recognition results for the given user.
func FindByUserID(userID string) []model.RecognitionResult {
	return memoryStore[userID]
}

// Reset clears the memory store — used for testing.
func Reset() {
	memoryStore = make(map[string][]model.RecognitionResult)
}
