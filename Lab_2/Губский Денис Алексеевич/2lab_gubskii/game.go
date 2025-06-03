package main

const gridSize = 4

// Game: только данные состояния игры
type Game struct {
	Grid        [gridSize][gridSize]int `json:"grid"`
	Score       int                     `json:"score"`
	Over        bool                    `json:"over"`
	NewTiles    [][2]int                `json:"newTiles"`
	MergedTiles [][2]int                `json:"mergedTiles"`
	Win         bool                    `json:"win"`
}
