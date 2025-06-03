/*
Project: Контрольная работа. Алгоритмы и Go Routines.
Description: Использует Go Routines и каналы для генерации случайных чисел и обработки их квадратных корней.
Author: Никита Мякишев (Вариант 3)
License: GPLv3
History:
  - [03.04.2025 17:48]: Initial task2
*/

package task2

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Генерирует случайные числа от 1 до 100 и отправляет их в канал
func generateNumbers(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	for i := 0; i < 20; i++ {
		num := rand.Intn(100) + 1
		ch <- num
	}
	close(ch) // Закрываем канал, когда генерация чисел завершена
}

// Вычисляет квадратный корень числа и выводит, если результат — целое число
func processNumbers(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		sqrt := math.Sqrt(float64(num))
		if sqrt == math.Floor(sqrt) { // Проверка, является ли корень целым числом
			fmt.Printf("Число: %d, Квадратный корень: %.0f\n", num, sqrt)
		}
	}
}

func RunTask2() {
	// Создаем канал для передачи чисел
	ch := make(chan int)
	var wg sync.WaitGroup

	// Запускаем горутины
	wg.Add(2) // Количество горутин
	go generateNumbers(ch, &wg)
	go processNumbers(ch, &wg)

	// Ожидаем завершения всех горутин
	wg.Wait()
}
