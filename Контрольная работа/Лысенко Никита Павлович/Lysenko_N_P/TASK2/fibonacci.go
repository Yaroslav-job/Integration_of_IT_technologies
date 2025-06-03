/*
Package main предоставляет функциональность для генерации и отображения
последовательности чисел Фибоначчи.
Название: Fibonacci Generator
Краткое описание: Модуль для генерации последовательности чисел Фибоначчи.
Разработчик: Nikita Lysenko
Лицензия: GPLv3
История изменений: 1.0 - первая версия
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateFibonacci генерирует последовательность чисел Фибоначчи и отправляет ее в указанный канал.
// Аргументы:
//   ch chan<- int: канал, в который будут отправлены числа Фибоначчи.
//   n int: количество чисел Фибоначчи для генерации.
func GenerateFibonacci(ch chan<- int, n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch)
}

// PrintFibonacci получает числа из канала и выводит их на экран,
// добавляя к каждому числу случайное число от 1 до 10.
// Аргументы:
//   ch <-chan int: канал, из которого будут получены числа Фибоначчи.
func PrintFibonacci(ch <-chan int) {
	for num := range ch {
		randomNum := rand.Intn(10) + 1 // Генерируем случайное число от 1 до 10
		fmt.Printf("%d + %d = %d\n", num, randomNum, num+randomNum)
	}
}

// main является точкой входа в программу. Он инициализирует генерацию
// последовательности Фибоначчи и выводит её на экран.
func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	go GenerateFibonacci(ch, 10) // Генерируем 10 чисел Фибоначчи
	PrintFibonacci(ch)            // Печатаем числа
}
