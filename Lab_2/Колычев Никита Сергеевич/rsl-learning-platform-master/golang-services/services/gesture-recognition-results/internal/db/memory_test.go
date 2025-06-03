package db

import (
	"testing"
	"time"

	"gesture-recognition-results/internal/model"
)

func TestSaveAndFindByUserID(t *testing.T) {
	Reset()

	result := model.RecognitionResult{
		ID:         "test-id-001",
		UserID:     "user123",
		Gesture:    "hello",
		Confidence: 0.9,
		Timestamp:  time.Now(),
	}

	Save(result)

	results := FindByUserID("user123")
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	got := results[0]
	if got.ID != result.ID {
		t.Errorf("expected ID %s, got %s", result.ID, got.ID)
	}
	if got.Gesture != result.Gesture {
		t.Errorf("expected Gesture %s, got %s", result.Gesture, got.Gesture)
	}
}

func TestFindByUserID_NoResults(t *testing.T) {
	Reset()

	results := FindByUserID("unknown_user")
	if results != nil && len(results) > 0 {
		t.Errorf("expected no results, got %v", results)
	}
}

func TestReset(t *testing.T) {
	Reset()

	result := model.RecognitionResult{
		ID:         "reset-id",
		UserID:     "reset-user",
		Gesture:    "wave",
		Confidence: 0.95,
		Timestamp:  time.Now(),
	}
	Save(result)

	if len(FindByUserID("reset-user")) != 1 {
		t.Fatalf("expected 1 result before reset")
	}

	Reset()

	if len(FindByUserID("reset-user")) != 0 {
		t.Fatalf("expected 0 results after reset")
	}
}
