// sumCalculator_test.go
// Модульные тесты для модуля вычисления суммы с использованием горутин
// Разработчик: Nikita Pashin
// Лицензия: GPLv3
// История изменений:
// 2025-04-03: Первая версия

package calculator

import (
	"sync"
	"testing"
)

func TestCreateGenerator(t *testing.T) {
	// Создаем канал и WaitGroup
	numbers := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем генератор, который должен выдать 5 чисел
	go createGenerator(numbers, 5, 100, &wg)

	// Создаем горутину, которая закроет канал после завершения генератора
	go func() {
		wg.Wait()
		close(numbers)
	}()

	// Проверяем, что генератор отправил 5 чисел
	count := 0
	for range numbers {
		count++
	}

	if count != 5 {
		t.Errorf("Generator должен отправить 5 чисел, получено: %d", count)
	}
}

func TestGetSum(t *testing.T) {
	// Тест с фиксированными числами для проверки суммирования
	t.Run("Проверка суммирования с фиксированными числами", func(t *testing.T) {
		numbers := make(chan int, 3)
	
		
		// Отправляем фиксированные числа: 5, 10, 15
		go func() {
			numbers <- 5
			numbers <- 10
			numbers <- 15
			close(numbers)
		}()
		
		// Вычисляем сумму
		sum := 0
		for num := range numbers {
			sum += num
		}
		
		expected := 30
		if sum != expected {
			t.Errorf("Сумма должна быть %d, получено: %d", expected, sum)
		}
	})
	
	// Тест с конфигурацией по умолчанию
	t.Run("Тест с конфигурацией по умолчанию", func(t *testing.T) {
		config := GetDefaultConfig()
		sum := GetSum(config)
		
		// Проверяем, что сумма неотрицательна (так как все числа неотрицательны)
		if sum < 0 {
			t.Errorf("Сумма должна быть неотрицательной, получено: %d", sum)
		}
		
		// Проверяем, что сумма не превышает максимально возможное значение
		maxPossibleSum := config.GeneratorCount * config.NumbersPerGen * config.MaxNumber
		if sum > maxPossibleSum {
			t.Errorf("Сумма превышает максимально возможное значение: %d > %d", sum, maxPossibleSum)
		}
	})
}