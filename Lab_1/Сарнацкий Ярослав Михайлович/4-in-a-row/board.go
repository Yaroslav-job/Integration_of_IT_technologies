package main

import (
	"fmt"
)

var boardArr [6][7]string

func board_start() {
	for i := 0; i < 6; i++ {
		for k := 0; k < 7; k++ {
			boardArr[i][k] = "[  ]"
		}
	}
	board_print()
}

func board_print() {
	clear()
	for i := 0; i < 6; i++ {
		for k := 0; k < 7; k++ {
			fmt.Printf("%s", boardArr[i][k])
		}
		fmt.Print("\n")
	}
}
