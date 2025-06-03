package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gesture-recognition-results/internal/db"
	"gesture-recognition-results/internal/model"
)

// Clears in-memory DB before each test
func resetDB() {
	db.Reset()
}

func TestHandleRecognition_ValidPayload(t *testing.T) {
	resetDB()

	payload := `{
		"user_id": "user42",
		"gesture": "hello",
		"confidence": 0.95
	}`
	req := httptest.NewRequest(http.MethodPost, "/recognitions", bytes.NewBufferString(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	HandleRecognition(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}

	var result model.RecognitionResult
	if err := json.NewDecoder(rr.Body).Decode(&result); err != nil {
		t.Fatal("response JSON decode failed")
	}

	if result.UserID != "user42" {
		t.Errorf("expected user_id to be 'user42', got '%s'", result.UserID)
	}
	if result.Gesture != "hello" {
		t.Errorf("expected gesture 'hello', got '%s'", result.Gesture)
	}
}

func TestGetRecognitionsByUser(t *testing.T) {
	resetDB()

	// Pre-seed the DB
	db.Save(model.RecognitionResult{
		ID:         "abc",
		UserID:     "user42",
		Gesture:    "wave",
		Confidence: 0.88,
	})

	req := httptest.NewRequest(http.MethodGet, "/recognitions/user42", nil)
	rr := httptest.NewRecorder()

	GetRecognitionsByUser(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var results []model.RecognitionResult
	if err := json.NewDecoder(rr.Body).Decode(&results); err != nil {
		t.Fatal("response JSON decode failed")
	}
	if len(results) != 1 || results[0].Gesture != "wave" {
		t.Errorf("expected 1 result with gesture 'wave', got %v", results)
	}
}

func TestHandleRecognition_InvalidConfidence(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/recognitions", strings.NewReader(`{
		"user_id": "x", "gesture": "bad", "confidence": 2.0
	}`))
	rr := httptest.NewRecorder()

	HandleRecognition(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for invalid confidence, got %d", rr.Code)
	}
}
