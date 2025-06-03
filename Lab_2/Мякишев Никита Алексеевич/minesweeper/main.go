package main

import (
	"encoding/json"
	"minesweeper/game"
	"net/http"
)

var board *game.Board

func main() {
	board = game.NewBoard(10, 10, 15)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/board", getBoard)
	http.HandleFunc("/api/reveal", revealCell)
	http.HandleFunc("/api/new", newGame)
	http.HandleFunc("/api/flag", toggleFlag)

	http.ListenAndServe(":8080", nil)
}

// Переключение флага
func toggleFlag(w http.ResponseWriter, r *http.Request) {
	var pos struct{ X, Y int }
	json.NewDecoder(r.Body).Decode(&pos)
	board.ToggleFlag(pos.X, pos.Y)
	w.WriteHeader(http.StatusOK)
}

// Отдаёт текущее состояние поля
func getBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(board.Grid)
}

// Открывает ячейку по координатам
func revealCell(w http.ResponseWriter, r *http.Request) {
	var pos struct{ X, Y int }
	json.NewDecoder(r.Body).Decode(&pos)
	board.Reveal(pos.X, pos.Y)
	w.WriteHeader(http.StatusOK)
}

// Новая игра
func newGame(w http.ResponseWriter, r *http.Request) {
	board.Reset()
	w.WriteHeader(http.StatusOK)
}
