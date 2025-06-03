# Sudoku Plus_v2

## How to Run

### Terminal Version
- go run main.go

### Web Version
- go run main.go -web -port 5000
- Then open your browser to http://localhost:5000

### Controls

#### Terminal:
- Arrow Keys: Move cursor between cells
- Numbers 1-9: Enter a number in the selected cell
- 0, Delete, or Backspace: Clear a cell
- ESC: Access the pause menu

#### Web:
- Click on a cell to select it
- Click on a number from the number pad to enter it in the selected cell
- Use the "Clear" button to remove a number from a cell

### Game Rules
1. Each row must contain all digits from 1 to 9
2. Each column must contain all digits from 1 to 9
3. Each 3×3 subgrid must contain all digits from 1 to 9.