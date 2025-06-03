package model

import "time"

// RecognitionResult represents a single gesture recognition result.
// Used for storing gesture, confidence level, and user tracking.
type RecognitionResult struct {
	ID         string    `json:"id"`         // Server-generated unique result ID
	UserID     string    `json:"user_id"`    // ID of the user performing the gesture
	Gesture    string    `json:"gesture"`    // Recognized gesture label
	Confidence float32   `json:"confidence"` // ML model confidence score (0.0 - 1.0)
	Timestamp  time.Time `json:"timestamp"`  // Server-side UTC timestamp
}
