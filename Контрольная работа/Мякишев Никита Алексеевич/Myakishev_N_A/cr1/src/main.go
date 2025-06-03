/*
Project: Контрольная работа. Алгоритмы и Go Routines.
Description:
	1. Находит наибольший общий префикс в заданном массиве строк.
  	2. Использует Go Routines и каналы для генерации случайных чисел и обработки их квадратных корней.
Author: Никита Мякишев (Вариант 3)
License: GPLv3
History:
  - [03.04.2025 17:35]: Initial main
  - [03.04.2025 18:04]: add task2
*/

package main

import (
	"cr1/src/task1"
	"cr1/src/task2"
	"fmt"
)

func main() {
	// Задача 1: Найти наибольший общий префикс в заданном массиве строк
	words := []string{"flower", "flow", "flight"}
	fmt.Println("Наибольший общий префикс:", task1.LongestCommonPrefix(words))

	task2.RunTask2()
}
