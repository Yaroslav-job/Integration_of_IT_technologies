// sumCalculator.go
// Модуль для вычисления суммы случайных чисел с использованием горутин и каналов
// Разработчик: Nikita Pashin
// Лицензия: GPLv3
// История изменений:
// 2025-04-03: Первая версия

package calculator

import (
	"math/rand"
	"sync"
	"time"
)

// Config содержит настройки для генерации и обработки чисел
type Config struct {
	GeneratorCount int // Количество генераторов
	NumbersPerGen  int // Количество чисел на один генератор
	MaxNumber      int // Максимальное генерируемое число
}

// GetDefaultConfig возвращает конфигурацию по умолчанию
func GetDefaultConfig() Config {
	return Config{
		GeneratorCount: 5,
		NumbersPerGen:  20,
		MaxNumber:      100,
	}
}

// createGenerator создает горутину, которая генерирует случайные числа и отправляет их в канал
func createGenerator(numbers chan<- int, count int, maxNumber int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		num := r.Intn(maxNumber)
		numbers <- num
	}
}

// getSum запускает горутины для генерации чисел и вычисления их суммы
// Возвращает общую сумму сгенерированных чисел
func GetSum(config Config) int {
	// Инициализация канала и WaitGroup
	numbers := make(chan int, config.GeneratorCount)
	var wg sync.WaitGroup

	// Запуск генераторов
	wg.Add(config.GeneratorCount)
	for i := 0; i < config.GeneratorCount; i++ {
		go createGenerator(numbers, config.NumbersPerGen, config.MaxNumber, &wg)
	}

	// Горутина для закрытия канала после завершения всех генераторов
	go func() {
		wg.Wait()
		close(numbers)
	}()

	// Вычисление суммы в основной горутине
	sum := 0
	for num := range numbers {
		sum += num
	}

	return sum
}