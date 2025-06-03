// fibonacci_test.go
// Название: Fibonacci Tests
// Краткое описание: Модульные тесты для проверки генерации чисел Фибоначчи.
// Разработчик: Ваше Имя
// Лицензия: GPLv3
// История изменений: 1.0 - первая версия

package main

import (
  "testing"
)

func TestGenerateFibonacci(t *testing.T) {
  ch := make(chan int)
  go GenerateFibonacci(ch, 10)

  count := 0
  for num := range ch {
    if count == 0 && num != 0 {
      t.Errorf("Expected first Fibonacci number to be 0, got %d", num)
    }
    if count == 1 && num != 1 {
      t.Errorf("Expected second Fibonacci number to be 1, got %d", num)
    }
    count++
  }

  if count != 10 {
    t.Errorf("Expected to receive 10 Fibonacci numbers, got %d", count)
  }
}
