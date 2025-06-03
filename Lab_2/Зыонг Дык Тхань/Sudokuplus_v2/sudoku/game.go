package sudoku

import (
	"errors"
)

const (
	GridSize    = 9
	SubGridSize = 3
	MaxLevels   = 10
)

// Game represents a Sudoku game state
type Game struct {
	ID           string                  `json:"id"`
	Grid         [GridSize][GridSize]int `json:"grid"`
	Solution     [GridSize][GridSize]int `json:"solution,omitempty"`
	OriginalGrid [GridSize][GridSize]int `json:"originalGrid"`
	Level        int                     `json:"level"`
	Completed    bool                    `json:"completed"`
}

// Create a new game
func NewGame(level int) (*Game, error) {
	if level < 1 || level > MaxLevels {
		return nil, errors.New("invalid level number")
	}

	game := &Game{
		Level:     level,
		Completed: false,
	}

	// Generate a new puzzle
	fullGrid, err := GenerateFullGrid()
	if err != nil {
		return nil, err
	}

	// Copy the full grid as the solution
	game.Solution = fullGrid
	cellsToRemove := 30 + (level * 2)
	game.Grid = RemoveNumbersByCount(fullGrid, cellsToRemove)
	game.OriginalGrid = game.Grid

	return game, nil
}

// IsValidMove checks if placing a number at a position is valid
func (g *Game) IsValidMove(row, col, num int) bool {
	// Check if the cell is an original cell (can't modify)
	if g.OriginalGrid[row][col] != 0 {
		return false
	}

	// Check row
	for i := 0; i < GridSize; i++ {
		if g.Grid[row][i] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < GridSize; i++ {
		if g.Grid[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	boxRow, boxCol := (row/SubGridSize)*SubGridSize, (col/SubGridSize)*SubGridSize
	for i := 0; i < SubGridSize; i++ {
		for j := 0; j < SubGridSize; j++ {
			if g.Grid[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}

// MakeMove places a number at the specified position
func (g *Game) MakeMove(row, col, num int) bool {

	if row < 0 || row >= GridSize || col < 0 || col >= GridSize {
		return false
	}

	if g.OriginalGrid[row][col] != 0 {
		return false
	}

	if num == 0 {
		g.Grid[row][col] = 0
		return true
	}

	if !g.IsValidMove(row, col, num) {
		return false
	}

	g.Grid[row][col] = num
	if g.IsCompleted() {
		g.Completed = true
	}

	return true
}

func (g *Game) IsCompleted() bool {
	// Check if all cells are filled
	for row := 0; row < GridSize; row++ {
		for col := 0; col < GridSize; col++ {
			if g.Grid[row][col] == 0 {
				return false
			}
		}
	}

	// Check if each row contains all numbers 1-9
	for row := 0; row < GridSize; row++ {
		seen := [GridSize + 1]bool{}
		for col := 0; col < GridSize; col++ {
			num := g.Grid[row][col]
			if num < 1 || num > GridSize || seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	// Check if each column contains all numbers 1-9
	for col := 0; col < GridSize; col++ {
		seen := [GridSize + 1]bool{}
		for row := 0; row < GridSize; row++ {
			num := g.Grid[row][col]
			if num < 1 || num > GridSize || seen[num] {
				return false
			}
			seen[num] = true
		}
	}

	// Check if each 3x3 box contains all numbers 1-9
	for boxRow := 0; boxRow < GridSize; boxRow += SubGridSize {
		for boxCol := 0; boxCol < GridSize; boxCol += SubGridSize {
			seen := [GridSize + 1]bool{}
			for i := 0; i < SubGridSize; i++ {
				for j := 0; j < SubGridSize; j++ {
					num := g.Grid[boxRow+i][boxCol+j]
					if num < 1 || num > GridSize || seen[num] {
						return false
					}
					seen[num] = true
				}
			}
		}
	}

	return true
}

func (g *Game) ResetGame() {
	g.Grid = g.OriginalGrid
	g.Completed = false
}
