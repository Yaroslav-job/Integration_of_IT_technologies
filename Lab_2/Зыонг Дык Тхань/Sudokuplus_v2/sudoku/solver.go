package sudoku

// Solver handles solving Sudoku puzzles
type Solver struct {
	grid *[GridSize][GridSize]int
}

// Creates a new Sudoku solver
func NewSolver(grid *[GridSize][GridSize]int) *Solver {
	return &Solver{grid: grid}
}

// Solve the Sudoku puzzle
func (s *Solver) Solve() bool {
	return s.solveRecursive()
}

// solveRecursive to solve the puzzle
func (s *Solver) solveRecursive() bool {
	row, col, found := s.findEmptyCell()
	if !found {
		return true
	}

	for num := 1; num <= GridSize; num++ {
		if s.isValidPlacement(row, col, num) {
			s.grid[row][col] = num
			if s.solveRecursive() {
				return true
			}
			s.grid[row][col] = 0
		}
	}
	return false
}

// Finds the next empty cell
func (s *Solver) findEmptyCell() (int, int, bool) {
	for row := 0; row < GridSize; row++ {
		for col := 0; col < GridSize; col++ {
			if s.grid[row][col] == 0 {
				return row, col, true
			}
		}
	}
	return -1, -1, false
}

// Checks if placing a number is valid
func (s *Solver) isValidPlacement(row, col, num int) bool {
	// Check row
	for i := 0; i < GridSize; i++ {
		if s.grid[row][i] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < GridSize; i++ {
		if s.grid[i][col] == num {
			return false
		}
	}

	// Check 3x3 box
	boxRow, boxCol := (row/SubGridSize)*SubGridSize, (col/SubGridSize)*SubGridSize
	for i := 0; i < SubGridSize; i++ {
		for j := 0; j < SubGridSize; j++ {
			if s.grid[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}
