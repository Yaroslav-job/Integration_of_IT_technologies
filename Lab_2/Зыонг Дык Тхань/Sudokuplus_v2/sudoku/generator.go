package sudoku

import (
	"errors"
	"math/rand"
)

// Creates a fully solved Sudoku grid
func GenerateFullGrid() ([GridSize][GridSize]int, error) {
	var grid [GridSize][GridSize]int

	solver := NewSolver(&grid)
	if !solver.Solve() {
		return grid, errors.New("failed to generate valid Sudoku grid")
	}

	shuffleGrid(&grid)

	return grid, nil
}

// shuffleGrid applies transformations to create a different valid Sudoku grid
func shuffleGrid(grid *[GridSize][GridSize]int) {
	for i := 0; i < 10; i++ {
		switch rand.Intn(5) {
		case 0:
			// Swap two rows within the same block
			blockRow := rand.Intn(SubGridSize)
			row1 := blockRow*SubGridSize + rand.Intn(SubGridSize)
			row2 := blockRow*SubGridSize + rand.Intn(SubGridSize)
			if row1 != row2 {
				swapRows(grid, row1, row2)
			}
		case 1:
			// Swap two columns within the same block
			blockCol := rand.Intn(SubGridSize)
			col1 := blockCol*SubGridSize + rand.Intn(SubGridSize)
			col2 := blockCol*SubGridSize + rand.Intn(SubGridSize)
			if col1 != col2 {
				swapColumns(grid, col1, col2)
			}
		case 2:
			// Swap two row blocks
			block1 := rand.Intn(SubGridSize)
			block2 := rand.Intn(SubGridSize)
			if block1 != block2 {
				swapRowBlocks(grid, block1, block2)
			}
		case 3:
			// Swap two column blocks
			block1 := rand.Intn(SubGridSize)
			block2 := rand.Intn(SubGridSize)
			if block1 != block2 {
				swapColumnBlocks(grid, block1, block2)
			}
		case 4:
			transposeGrid(grid)
		}
	}
}

// swapRows swaps two rows in the grid
func swapRows(grid *[GridSize][GridSize]int, row1, row2 int) {
	grid[row1], grid[row2] = grid[row2], grid[row1]
}

// swapColumns swaps two columns in the grid
func swapColumns(grid *[GridSize][GridSize]int, col1, col2 int) {
	for row := 0; row < GridSize; row++ {
		grid[row][col1], grid[row][col2] = grid[row][col2], grid[row][col1]
	}
}

// swapRowBlocks swaps two row blocks
func swapRowBlocks(grid *[GridSize][GridSize]int, block1, block2 int) {
	for i := 0; i < SubGridSize; i++ {
		row1 := block1*SubGridSize + i
		row2 := block2*SubGridSize + i
		swapRows(grid, row1, row2)
	}
}

// swapColumnBlocks swaps two column blocks
func swapColumnBlocks(grid *[GridSize][GridSize]int, block1, block2 int) {
	for i := 0; i < SubGridSize; i++ {
		col1 := block1*SubGridSize + i
		col2 := block2*SubGridSize + i
		swapColumns(grid, col1, col2)
	}
}

// transposeGrid transposes the grid (rows become columns and vice versa)
func transposeGrid(grid *[GridSize][GridSize]int) {
	for i := 0; i < GridSize; i++ {
		for j := i + 1; j < GridSize; j++ {
			grid[i][j], grid[j][i] = grid[j][i], grid[i][j]
		}
	}
}

// Remove a specific number of cells from the grid
func RemoveNumbersByCount(fullGrid [GridSize][GridSize]int, cellsToRemove int) [GridSize][GridSize]int {
	grid := fullGrid

	positions := make([][2]int, GridSize*GridSize)
	idx := 0
	for i := 0; i < GridSize; i++ {
		for j := 0; j < GridSize; j++ {
			positions[idx] = [2]int{i, j}
			idx++
		}
	}

	rand.Shuffle(len(positions), func(i, j int) {
		positions[i], positions[j] = positions[j], positions[i]
	})

	removedCount := 0
	for _, pos := range positions {
		row, col := pos[0], pos[1]

		value := grid[row][col]

		grid[row][col] = 0

		if removedCount < cellsToRemove && isSolvable(grid) {
			removedCount++
		} else {
			grid[row][col] = value
		}

		if removedCount >= cellsToRemove {
			break
		}
	}

	return grid
}

func isSolvable(grid [GridSize][GridSize]int) bool {
	testGrid := grid
	solver := NewSolver(&testGrid)
	return solver.Solve()
}
