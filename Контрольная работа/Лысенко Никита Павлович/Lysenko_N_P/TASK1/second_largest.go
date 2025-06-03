/*
Package main предоставляет функции для работы с массивами целых чисел.
Название: Second Largest
Краткое описание: Поиск второго по величине числа массива.
Разработчик: Nikita Lysenko
Лицензия: GPLv3
История изменений: 1.0 - первая версия
*/
package main

import (
	"fmt"
	"math"
)

// FindSecondLargest находит второй по величине элемент в массиве целых чисел.
// Если массив содержит менее двух элементов, возвращается ошибка.
// Если второй по величине элемент не найден, также возвращается ошибка.
//
// Аргументы:
//   arr []int: массив целых чисел, в котором нужно найти второй по величине элемент.
//
// Возвращаемые значения:
//   int: второй по величине элемент в массиве.
//   error: ошибка, если массив содержит менее двух элементов или если второй по величине элемент не найден.
func FindSecondLargest(arr []int) (int, error) {
	if len(arr) < 2 {
		return 0, fmt.Errorf("массив должен содержать как минимум два элемента")
	}

	max := math.MinInt
	secondMax := math.MinInt

	for _, num := range arr {
		if num > max {
			secondMax = max
			max = num
		} else if num > secondMax && num < max {
			secondMax = num
		}
	}

	if secondMax == math.MinInt {
		return 0, fmt.Errorf("второй по величине элемент не найден")
	}

	return secondMax, nil
}

// main является точкой входа в программу. Он демонстрирует использование функции FindSecondLargest
// с примером массива целых чисел и выводит второй по величине элемент.
func main() {
	arr := []int{5, 2, 8, 1, 9, 4}
	second, err := FindSecondLargest(arr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Второй по величине элемент:", second)
	}
}
