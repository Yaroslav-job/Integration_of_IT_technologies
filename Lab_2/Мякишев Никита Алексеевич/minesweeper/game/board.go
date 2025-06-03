package game

import (
	"math/rand"
	"time"
)

type Board struct {
	Width, Height int
	Mines         int
	Grid          [][]*Cell
}

// Создать новое поле
func NewBoard(width, height, mines int) *Board {
	b := &Board{Width: width, Height: height, Mines: mines}
	b.initGrid()
	b.placeMines()
	b.calculateNeighbors()
	return b
}

// Инициализация пустого поля
func (b *Board) initGrid() {
	b.Grid = make([][]*Cell, b.Height)
	for y := 0; y < b.Height; y++ {
		b.Grid[y] = make([]*Cell, b.Width)
		for x := 0; x < b.Width; x++ {
			b.Grid[y][x] = &Cell{}
		}
	}
}

// Случайное размещение мин
func (b *Board) placeMines() {
	rand.Seed(time.Now().UnixNano())
	placed := 0
	for placed < b.Mines {
		x := rand.Intn(b.Width)
		y := rand.Intn(b.Height)
		if !b.Grid[y][x].IsMine {
			b.Grid[y][x].IsMine = true
			placed++
		}
	}
}

// Подсчёт мин вокруг каждой ячейки
func (b *Board) calculateNeighbors() {
	dirs := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1},
	}
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			count := 0
			for _, d := range dirs {
				nx, ny := x+d.dx, y+d.dy
				if nx >= 0 && ny >= 0 && nx < b.Width && ny < b.Height && b.Grid[ny][nx].IsMine {
					count++
				}
			}
			b.Grid[y][x].NeighborMines = count
		}
	}
}

// Открыть ячейку (рекурсивно если 0 мин рядом)
func (b *Board) Reveal(x, y int) {
	if x < 0 || y < 0 || x >= b.Width || y >= b.Height {
		return
	}
	cell := b.Grid[y][x]
	if cell.IsRevealed || cell.IsFlagged {
		return
	}
	cell.IsRevealed = true
	if cell.NeighborMines == 0 && !cell.IsMine {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx != 0 || dy != 0 {
					b.Reveal(x+dx, y+dy)
				}
			}
		}
	}
}

// Переключить флаг
func (b *Board) ToggleFlag(x, y int) {
	if x < 0 || y < 0 || x >= b.Width || y >= b.Height {
		return
	}
	cell := b.Grid[y][x]
	if cell.IsRevealed {
		return
	}
	cell.IsFlagged = !cell.IsFlagged
}

// Подсчёт флагов
func (b *Board) CountFlags() int {
	count := 0
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Grid[y][x].IsFlagged {
				count++
			}
		}
	}
	return count
}

// Сброс поля
func (b *Board) Reset() {
	b.initGrid()
	b.placeMines()
	b.calculateNeighbors()
}
