package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sudokuplus/sudoku"
)

// Game store to keep track of active games
type GameStore struct {
	ActiveGames map[string]*sudoku.Game
}

// Global game store
var gameStore = GameStore{
	ActiveGames: make(map[string]*sudoku.Game),
}

// JSON response helper function
func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Error response helper function
func writeError(w http.ResponseWriter, message string, status int) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	writeJSON(w, errorResponse{Error: message}, status)
}

// Creates a new game
func handleNewGame(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse level from request
	levelStr := r.FormValue("level")
	level, err := strconv.Atoi(levelStr)
	if err != nil || level < 1 || level > sudoku.MaxLevels {
		writeError(w, fmt.Sprintf("Invalid level. Expected 1-%d", sudoku.MaxLevels), http.StatusBadRequest)
		return
	}

	// Create a new game
	game, err := sudoku.NewGame(level)
	if err != nil {
		writeError(w, fmt.Sprintf("Failed to create new game: %v", err), http.StatusInternalServerError)
		return
	}
	gameID := fmt.Sprintf("%d", len(gameStore.ActiveGames)+1)
	game.ID = gameID
	gameStore.ActiveGames[gameID] = game
	writeJSON(w, game, http.StatusOK)
}

// Move in the game
func handleMove(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request parameters
	gameID := r.FormValue("gameId")
	rowStr := r.FormValue("row")
	colStr := r.FormValue("col")
	numStr := r.FormValue("num")

	// Validate parameters
	if gameID == "" {
		writeError(w, "Game ID is required", http.StatusBadRequest)
		return
	}

	// Get the game from store
	game, exists := gameStore.ActiveGames[gameID]
	if !exists {
		writeError(w, "Game not found", http.StatusNotFound)
		return
	}

	// Parse row, col, and num
	row, err1 := strconv.Atoi(rowStr)
	col, err2 := strconv.Atoi(colStr)
	num, err3 := strconv.Atoi(numStr)

	if err1 != nil || err2 != nil || err3 != nil {
		writeError(w, "Invalid row, column, or number", http.StatusBadRequest)
		return
	}

	// Make the move
	success := game.MakeMove(row, col, num)

	// Return the result
	result := map[string]interface{}{
		"success":   success,
		"completed": game.Completed,
	}

	writeJSON(w, result, http.StatusOK)
}
