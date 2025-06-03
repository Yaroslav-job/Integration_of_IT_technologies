package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
)

const gridSize = 9

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
	0: "\033[90m", // Grey empty cells
}

const (
	resetColor     = "\033[0m"
	highlightColor = "\033[47;30m"
)

type Sudoku struct {
	grid [gridSize][gridSize]int
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (s *Sudoku) printBoard(selectedRow, selectedCol, level int) {
	clearTerminal()
	fmt.Printf("Level: %d\n\n", level)
	for i := 0; i < gridSize; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j := 0; j < gridSize; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			color := colors[s.grid[i][j]]
			if i == selectedRow && j == selectedCol {
				fmt.Print(highlightColor, s.grid[i][j], " ", resetColor)
			} else {
				if s.grid[i][j] == 0 {
					fmt.Print(color, "■ ", resetColor)
				} else {
					fmt.Print(color, s.grid[i][j], " ", resetColor)
				}
			}
		}
		fmt.Println()
	}
}

func (s *Sudoku) isValidMove(row, col, num int) bool {
	for i := 0; i < gridSize; i++ {
		if s.grid[row][i] == num || s.grid[i][col] == num {
			return false
		}
	}
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s.grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) makeMove(row, col, num int) bool {
	if row >= 0 && row < gridSize && col >= 0 && col < gridSize && s.grid[row][col] == 0 {
		if s.isValidMove(row, col, num) {
			s.grid[row][col] = num
			return true
		}
	}
	return false
}

func generateSudoku() Sudoku {
	rand.Seed(time.Now().UnixNano())
	s := Sudoku{}
	for i := 0; i < 10; i++ {
		row, col, num := rand.Intn(gridSize), rand.Intn(gridSize), rand.Intn(9)+1
		if s.isValidMove(row, col, num) {
			s.grid[row][col] = num
		}
	}
	return s
}

func main() {
	sudoku := generateSudoku()
	level := 1
	fmt.Println("Use arrow keys to move, numbers 1-9 to input, and 'q' to quit.")
	sudoku.printBoard(0, 0, level)

	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to initialize keyboard input:", err)
		os.Exit(1)
	}
	defer keyboard.Close()

	row, col := 0, 0
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key input:", err)
			continue
		}

		if key == keyboard.KeyEsc || char == 'q' {
			fmt.Println("Exiting game...")
			os.Exit(0)
		} else if key == keyboard.KeyArrowUp && row > 0 {
			row--
		} else if key == keyboard.KeyArrowDown && row < gridSize-1 {
			row++
		} else if key == keyboard.KeyArrowLeft && col > 0 {
			col--
		} else if key == keyboard.KeyArrowRight && col < gridSize-1 {
			col++
		} else if char >= '1' && char <= '9' {
			num := int(char - '0')
			if sudoku.makeMove(row, col, num) {
				sudoku.printBoard(row, col, level)
			} else {
				fmt.Println("Invalid move.")
			}
		}
		sudoku.printBoard(row, col, level)
	}
}
