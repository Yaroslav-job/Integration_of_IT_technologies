package server

type GameState struct {
	Board    [6][7]string `json:"board"`
	Turn     string       `json:"turn"`
	GameOver bool         `json:"game_over"`
	Winner   string       `json:"winner"`
}

func ResetGame() {
	game = GameState{
		Turn:  "🟢",
		Board: [6][7]string{},
	}
}

func checkWin(marker string, b [6][7]string) bool {
	// Проверка на победу
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if j+3 < 7 &&
				b[i][j] == marker &&
				b[i][j+1] == marker &&
				b[i][j+2] == marker &&
				b[i][j+3] == marker {
				return true
			}
			if i+3 < 6 &&
				b[i][j] == marker &&
				b[i+1][j] == marker &&
				b[i+2][j] == marker &&
				b[i+3][j] == marker {
				return true
			}
			if i+3 < 6 && j+3 < 7 &&
				b[i][j] == marker &&
				b[i+1][j+1] == marker &&
				b[i+2][j+2] == marker &&
				b[i+3][j+3] == marker {
				return true
			}
			if i-3 >= 0 && j+3 < 7 &&
				b[i][j] == marker &&
				b[i-1][j+1] == marker &&
				b[i-2][j+2] == marker &&
				b[i-3][j+3] == marker {
				return true
			}
		}
	}
	return false
}

func isDraw(b [6][7]string) bool {
	// Проверка на ничью
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			if b[i][j] == "" {
				return false
			}
		}
	}
	return true
}
