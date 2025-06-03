package server

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	game     GameState
	mu       sync.Mutex
	gameType string // Тип игры: "player" или "bot"
)

func StartGame(w http.ResponseWriter, r *http.Request) {
	gameType := r.URL.Query().Get("type")
	if gameType == "bot" {
		// Инициализируем игру против бота
		game.GameOver = false
		game.Turn = "🟢"
		// Возможно, добавить логику для бота (если нужно)
	} else {
		// Инициализируем игру "Игрок против Игрока"
		game.GameOver = false
		game.Turn = "🟢"
	}
	json.NewEncoder(w).Encode(game)
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	if game.GameOver {
		json.NewEncoder(w).Encode(game)
		return
	}

	colStr := r.URL.Query().Get("column")
	col, err := strconv.Atoi(colStr)
	if err != nil || col < 0 || col > 6 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Проверка, что фишка еще не поставлена в выбранной колонке
	for i := 5; i >= 0; i-- {
		if game.Board[i][col] == "" {
			game.Board[i][col] = game.Turn
			if checkWin(game.Turn, game.Board) {
				game.GameOver = true
				game.Winner = game.Turn
			} else {
				if isDraw(game.Board) {
					game.GameOver = true
				}
				// Переключаем ход
				switch game.Turn {
				case "🟢":
					game.Turn = "🔴"
				case "🔴":
					game.Turn = "🟢"
				}
			}
			break
		}
	}

	// Если игра против бота, делаем ход бота
	if game.GameOver == false && game.Turn == "🔴" && gameType == "bot" {
		makeSmartBotMove()
	}

	// Отправляем обновленное состояние игры
	json.NewEncoder(w).Encode(game)
}

func makeSmartBotMove() {
	// Сначала проверим, может ли бот выиграть или заблокировать игрока
	col := findWinningMove("🔴") // Попытаться выиграть
	if col == -1 {
		col = findWinningMove("🟢") // Если не может выиграть, попробовать заблокировать игрока
	}

	if col == -1 {
		// Если нет выигрыша или блокировки, делаем случайный ход
		col = getRandomMove()
	}

	// Выполним ход в найденную колонку
	for i := 5; i >= 0; i-- {
		if game.Board[i][col] == "" {
			game.Board[i][col] = "🔴"
			if checkWin("🔴", game.Board) {
				game.GameOver = true
				game.Winner = "🔴"
			} else if isDraw(game.Board) {
				game.GameOver = true
			} else {
				game.Turn = "🟢" // Ход игрока
			}
			return
		}
	}
}

func findWinningMove(marker string) int {
	// Проверяем каждую колонку, если есть возможность выиграть — возвращаем колонку
	for col := 0; col < 7; col++ {
		for row := 5; row >= 0; row-- {
			if game.Board[row][col] == "" {
				game.Board[row][col] = marker
				if checkWin(marker, game.Board) {
					game.Board[row][col] = "" // Снимаем фишку после проверки
					return col
				}
				game.Board[row][col] = "" // Снимаем фишку после проверки
				break
			}
		}
	}
	return -1
}

func getRandomMove() int {
	// Делаем случайный ход
	rand.Seed(time.Now().UnixNano())
	for {
		col := rand.Intn(7) // Случайная колонка
		for row := 5; row >= 0; row-- {
			if game.Board[row][col] == "" {
				return col
			}
		}
	}
}

func HandleStart(w http.ResponseWriter, r *http.Request) {
	// Получаем тип игры: "player" или "bot"
	gameType = r.URL.Query().Get("type")
	ResetGame()
	json.NewEncoder(w).Encode(game)
}

func HandleReset(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	ResetGame()
	json.NewEncoder(w).Encode(game)
}
