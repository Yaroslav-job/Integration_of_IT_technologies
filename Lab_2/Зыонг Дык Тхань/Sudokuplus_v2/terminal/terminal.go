package terminal

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sudokuplus/sudoku"
	"time"

	"github.com/eiannone/keyboard"
)

// Colors for the terminal UI
var colors = map[int]string{
	1: "\033[31m", // Red
	2: "\033[32m", // Green
	3: "\033[33m", // Yellow
	4: "\033[34m", // Blue
	5: "\033[35m", // Magenta
	6: "\033[36m", // Cyan
	7: "\033[91m", // Light Red
	8: "\033[92m", // Light Green
	9: "\033[93m", // Light Yellow
	0: "\033[90m", // Grey for empty cells
}

const (
	resetColor     = "\033[0m"
	originalColor  = "\033[1m"
	selectionColor = "\033[44;37m"
)

type TerminalGame struct {
	game        *sudoku.Game
	selectedRow int
	selectedCol int
}

func StartGame() {
	keyboard.Close()
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard input:", err)
		os.Exit(1)
	}

	game := &TerminalGame{
		selectedRow: 0,
		selectedCol: 0,
	}

	defer func() {
		fmt.Print("\033[?25h") // Show cursor when exiting
		keyboard.Close()
	}()

	game.showMainMenu()
}

func (t *TerminalGame) showMainMenu() {
	keyboard.Close()
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard input:", err)
		os.Exit(1)
	}
	fmt.Print("\033[?25h")

	for {
		clearTerminal()
		fmt.Print("\033[?25h")
		fmt.Println("===== SUDOKU PLUS =====")
		fmt.Println("1. New Game")
		fmt.Println("2. Quit")
		fmt.Print("\nSelect an option (1-2): ")

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch char {
		case '1':
			t.startNewGame(1)
		case '2':
			fmt.Println("\nThanks for playing! Goodbye.")
			keyboard.Close()
			os.Exit(0)
		}
	}
}

func (t *TerminalGame) startNewGame(difficulty int) {
	game, err := sudoku.NewGame(difficulty)
	if err != nil {
		fmt.Println("Error creating new game:", err)
		return
	}

	t.game = game
	t.selectedRow = 0
	t.selectedCol = 0

	keyboard.Close()
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard input:", err)
		os.Exit(1)
	}

	t.playGame()
}

func (t *TerminalGame) playGame() {
	keyboard.Close()
	time.Sleep(100 * time.Millisecond)
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard input:", err)
		os.Exit(1)
	}
	defer keyboard.Close()

	clearTerminal()

	fmt.Print("\033[?25l")

	fmt.Println("===== SUDOKU PLUS =====")
	fmt.Println("\nCONTROL INSTRUCTIONS:")
	fmt.Println("- Use arrow keys (↑, ↓, ←, →) to move between cells")
	fmt.Println("- Enter numbers 1-9 to fill the selected cell")
	fmt.Println("- Press 0, Delete or Backspace to clear a cell")
	fmt.Println("- Press ESC to access the menu")
	fmt.Println("\nPress any key to start...")

	for {
		_, _, err := keyboard.GetSingleKey()
		if err == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}

	var lastRenderedBoard string

	refreshInterval := 100 * time.Millisecond
	lastRefreshTime := time.Now()

	boardChanged := true

	for {
		if boardChanged && time.Since(lastRefreshTime) > refreshInterval {
			var buffer bytes.Buffer
			t.renderBoardToBuffer(&buffer)
			newBoard := buffer.String()

			if newBoard != lastRenderedBoard {
				clearTerminal()
				fmt.Print(newBoard)
				lastRenderedBoard = newBoard
				lastRefreshTime = time.Now()
			}
			boardChanged = false
		}

		// Get keyboard input
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key input:", err)
			if err.Error() == "keyboard not opened" {
				keyboard.Close()
				if err := keyboard.Open(); err != nil {
					fmt.Println("Failed to reinitialize keyboard:", err)
					os.Exit(1)
				}
			}
			continue
		}

		switch {
		case key == keyboard.KeyEsc:
			t.showPauseMenu()
			return
		case key == keyboard.KeyArrowUp:
			t.selectedRow = max(0, t.selectedRow-1)
			boardChanged = true
		case key == keyboard.KeyArrowDown:
			t.selectedRow = min(sudoku.GridSize-1, t.selectedRow+1)
			boardChanged = true
		case key == keyboard.KeyArrowLeft:
			t.selectedCol = max(0, t.selectedCol-1)
			boardChanged = true
		case key == keyboard.KeyArrowRight:
			t.selectedCol = min(sudoku.GridSize-1, t.selectedCol+1)
			boardChanged = true
		case char >= '1' && char <= '9':
			num := int(char - '0')
			// Check if the cell can be filled
			if t.game.OriginalGrid[t.selectedRow][t.selectedCol] == 0 {
				success := t.game.MakeMove(t.selectedRow, t.selectedCol, num)
				boardChanged = true
				if success && t.game.Completed {
					t.showCompletionScreen()
					return
				}
			}
		case char == '0' || key == keyboard.KeyBackspace || key == keyboard.KeyDelete:
			if t.game.OriginalGrid[t.selectedRow][t.selectedCol] == 0 {
				t.game.MakeMove(t.selectedRow, t.selectedCol, 0)
				boardChanged = true
			}

		}
	}
}

// Renders the game board
func (t *TerminalGame) renderBoardToBuffer(buffer *bytes.Buffer) {
	buffer.WriteString("===== SUDOKU PLUS =====\n")
	buffer.WriteString(fmt.Sprintf("Level: %d/%d | Press ESC for menu\n",
		t.game.Level, sudoku.MaxLevels))
	buffer.WriteString("Use arrow keys to navigate and numbers 1-9 to fill cells\n\n")

	// Draw the Sudoku grid
	buffer.WriteString("  ┌───┬───┬───┐───┬───┬───┐───┬───┬───┐\n")

	for i := 0; i < sudoku.GridSize; i++ {
		buffer.WriteString("  │")

		for j := 0; j < sudoku.GridSize; j++ {
			isSelected := (i == t.selectedRow && j == t.selectedCol)
			isOriginal := t.game.OriginalGrid[i][j] != 0

			cellValue := t.game.Grid[i][j]

			if isSelected {
				if cellValue == 0 {
					buffer.WriteString(fmt.Sprintf("%s □ %s", selectionColor, resetColor))
				} else {
					buffer.WriteString(fmt.Sprintf("%s %d %s", selectionColor, cellValue, resetColor))
				}
			} else if isOriginal {
				buffer.WriteString(fmt.Sprintf(" %s%s%d%s ", originalColor, colors[cellValue], cellValue, resetColor))
			} else if cellValue == 0 {
				// Empty cell
				buffer.WriteString(fmt.Sprintf(" %s□%s ", colors[0], resetColor))
			} else {
				buffer.WriteString(fmt.Sprintf(" %s%d%s ", colors[cellValue], cellValue, resetColor))
			}

			if j < sudoku.GridSize-1 {
				if (j+1)%3 == 0 {
					buffer.WriteString("│")
				} else {
					buffer.WriteString("│")
				}
			}
		}
		buffer.WriteString("│\n")

		if i < sudoku.GridSize-1 {
			if (i+1)%3 == 0 {
				buffer.WriteString("  ├───┼───┼───┤───┼───┼───┤───┼───┼───┤\n")
			} else {
				buffer.WriteString("  ├───┼───┼───┼───┼───┼───┼───┼───┼───┤\n")
			}
		}
	}

	buffer.WriteString("  └───┴───┴───┘───┴───┴───┘───┴───┴───┘\n")

	buffer.WriteString("\nControls:\n")
	buffer.WriteString("  Arrow keys: Move cursor   |   1-9: Enter number   |   0/Del/Backspace: Clear cell\n")
}

func (t *TerminalGame) showPauseMenu() {
	keyboard.Close()
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard input:", err)
		os.Exit(1)
	}

	fmt.Print("\033[?25h")

	for {
		clearTerminal()
		fmt.Print("\033[?25h")
		fmt.Println("===== GAME PAUSED =====")
		fmt.Println("1. Resume Game")
		fmt.Println("2. Reset Game")
		fmt.Println("3. Return to Main Menu")
		fmt.Print("\nSelect an option (1-3): ")

		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		switch char {
		case '1':
			// Resume game
			keyboard.Close()
			if err := keyboard.Open(); err != nil {
				fmt.Println("Failed to initialize keyboard input:", err)
				os.Exit(1)
			}
			t.playGame()
			return
		case '2':
			// Reset game
			t.game.ResetGame()
			keyboard.Close()
			if err := keyboard.Open(); err != nil {
				fmt.Println("Failed to initialize keyboard input:", err)
				os.Exit(1)
			}
			t.playGame()
			return
		case '3':
			// Return to main menu
			return
		}
	}
}

func (t *TerminalGame) showCompletionScreen() {
	keyboard.Close()
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard input:", err)
		os.Exit(1)
	}

	fmt.Print("\033[?25h")

	clearTerminal()
	fmt.Print("\033[?25h")

	fmt.Println("===== PUZZLE COMPLETED! =====")

	fmt.Printf("\nLevel %d Completed!\n", t.game.Level)

	if t.game.Level < sudoku.MaxLevels {
		fmt.Println("\nAdvancing to the next level...")
		time.Sleep(1500 * time.Millisecond)
		t.startNewGame(t.game.Level + 1)
		return
	} else {
		fmt.Println("\nCongratulations! You have completed all levels!")
		fmt.Println("\nPress any key to return to the main menu...")
		keyboard.GetSingleKey()
	}
}

// clearTerminal clears the terminal screen
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Print("\033[H\033[2J")

	fmt.Print("\033[?25l")
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
