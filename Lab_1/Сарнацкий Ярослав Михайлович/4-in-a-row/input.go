package main

import (
	"fmt"
)

func input(marker_type string) {
	var column int
	for {
		fmt.Printf("Введите номер колонки(от 1 до 7) куда хотите поставить фишку: ")
		fmt.Scan(&column)
		if column < 1 || column > 7 {
			fmt.Println("Введено недопустимое значение")
		} else {
			answer := check_progress(column-1, marker_type)
			if answer == 1 {
				break
			}
		}
	}
}

func bot(marker_type string) {
	for {
		column := random()
		if column >= 1 && column <= 7 {
			answer := check_progress(column-1, marker_type)
			if answer == 1 {
				break
			}
		}
	}
}

func check_progress(column int, marker_type string) int {
	var n int
	var s = 5

	for i := 5; i >= 0; i-- {
		if boardArr[i][column] != "[  ]" {
			n += 1
		}
	}

	s = s - n

	if n == 6 {
		fmt.Println("Колонка уже заполнена. Пожалуйста выберите другую колонку.")
	} else {
		boardArr[s][column] = marker_type
		return 1
	}
	return 0
}
