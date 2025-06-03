// main.go
// Основной файл программы, демонстрирующий работу реализованных модулей
// Разработчик: Nikita Pashin
// Лицензия: GPLv3
// История изменений:
// 2025-04-03: Первая версия

package main

import (
	"fmt"
	"goproject/calculator"
	"goproject/rotation"
)

func main() {
	// Демонстрация циклического сдвига массива
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный массив:", arr)
	
	n := 2
	rotation.RotateRight(arr, n)
	fmt.Printf("Массив после сдвига на %d позиции вправо: %v\n\n", n, arr)
	
	// Демонстрация вычисления суммы с использованием горутин
	config := calculator.GetDefaultConfig()
	fmt.Printf("Запуск %d генераторов, каждый генерирует %d чисел...\n", 
		config.GeneratorCount, config.NumbersPerGen)
	
	sum := calculator.GetSum(config)
	fmt.Printf("Общая сумма сгенерированных чисел: %d\n", sum)
}
