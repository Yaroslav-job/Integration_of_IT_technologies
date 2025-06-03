package main

import (
	"math/rand"
)

type GameLogic struct {
	Game *Game
}

func NewGameLogic() *GameLogic {
	game := &Game{}
	logic := &GameLogic{Game: game}
	logic.AddRandomTile()
	logic.AddRandomTile()
	return logic
}

func (gl *GameLogic) AddRandomTile() {
	emptyCells := []struct{ X, Y int }{}
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if gl.Game.Grid[x][y] == 0 {
				emptyCells = append(emptyCells, struct{ X, Y int }{x, y})
			}
		}
	}
	if len(emptyCells) > 0 {
		cell := emptyCells[rand.Intn(len(emptyCells))]
		gl.Game.Grid[cell.X][cell.Y] = 2
		gl.Game.NewTiles = append(gl.Game.NewTiles, [2]int{cell.X, cell.Y})
	}
}

func (gl *GameLogic) Move(direction string) {
	if gl.Game.Over || gl.Game.Win {
		return
	}
	gl.Game.NewTiles = nil
	gl.Game.MergedTiles = nil
	switch direction {
	case "up":
		for x := 0; x < gridSize; x++ {
			gl.mergeColumn(x, 0, 1)
		}
	case "down":
		for x := 0; x < gridSize; x++ {
			gl.mergeColumn(x, gridSize-1, -1)
		}
	case "left":
		for y := 0; y < gridSize; y++ {
			gl.mergeRow(y, 0, 1)
		}
	case "right":
		for y := 0; y < gridSize; y++ {
			gl.mergeRow(y, gridSize-1, -1)
		}
	}
	gl.AddRandomTile()
	gl.checkGameOver()
	gl.checkWin()
}

func (gl *GameLogic) mergeRow(y, start, step int) {
	merged := make([]bool, gridSize)
	for i := start; i >= 0 && i < gridSize; i += step {
		if gl.Game.Grid[i][y] == 0 {
			continue
		}
		for j := i - step; j >= 0 && j < gridSize; j -= step {
			if gl.Game.Grid[j][y] == 0 {
				gl.Game.Grid[j][y], gl.Game.Grid[j+step][y] = gl.Game.Grid[j+step][y], 0
			} else if gl.Game.Grid[j][y] == gl.Game.Grid[j+step][y] && !merged[j] {
				gl.Game.Grid[j][y] *= 2
				gl.Game.Score += gl.Game.Grid[j][y]
				gl.Game.Grid[j+step][y] = 0
				merged[j] = true
				gl.Game.MergedTiles = append(gl.Game.MergedTiles, [2]int{j, y})
				break
			} else {
				break
			}
		}
	}
}

func (gl *GameLogic) mergeColumn(x, start, step int) {
	merged := make([]bool, gridSize)
	for i := start; i >= 0 && i < gridSize; i += step {
		if gl.Game.Grid[x][i] == 0 {
			continue
		}
		for j := i - step; j >= 0 && j < gridSize; j -= step {
			if gl.Game.Grid[x][j] == 0 {
				gl.Game.Grid[x][j], gl.Game.Grid[x][j+step] = gl.Game.Grid[x][j+step], 0
			} else if gl.Game.Grid[x][j] == gl.Game.Grid[x][j+step] && !merged[j] {
				gl.Game.Grid[x][j] *= 2
				gl.Game.Score += gl.Game.Grid[x][j]
				gl.Game.Grid[x][j+step] = 0
				merged[j] = true
				gl.Game.MergedTiles = append(gl.Game.MergedTiles, [2]int{x, j})
				break
			} else {
				break
			}
		}
	}
}

func (gl *GameLogic) checkGameOver() {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if gl.Game.Grid[x][y] == 0 {
				return
			}
			if x > 0 && gl.Game.Grid[x][y] == gl.Game.Grid[x-1][y] {
				return
			}
			if y > 0 && gl.Game.Grid[x][y] == gl.Game.Grid[x][y-1] {
				return
			}
		}
	}
	gl.Game.Over = true
}

func (gl *GameLogic) checkWin() {
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if gl.Game.Grid[x][y] == 2048 {
				gl.Game.Win = true
				return
			}
		}
	}
}
