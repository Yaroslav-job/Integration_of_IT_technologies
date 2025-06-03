// second_largest_test.go
// Название: Second Largest Tests
// Краткое описание: Модульные тесты для проверки нахождения второго по величине элемента.
// Разработчик: Ваше Имя
// Лицензия: GPLv3
// История изменений: 1.0 - первая версия

package main

import (
  "testing"
)

func TestFindSecondLargest(t *testing.T) {
  tests := []struct {
    arr      []int
    expected int
  }{
    {[]int{5, 2, 8, 1, 9, 4}, 8},
    {[]int{1, 2, 3, 4, 5}, 4},
    {[]int{10, 10, 10, 10}, 0}, // Все элементы одинаковы
    {[]int{3}, 0},              // Массив содержит только один элемент
  }

  for _, test := range tests {
    got, err := FindSecondLargest(test.arr)
    if err != nil && test.expected != 0 {
      t.Errorf("FindSecondLargest(%v) returned an error: %v", test.arr, err)
    }
    if got != test.expected {
      t.Errorf("FindSecondLargest(%v) = %d; want %d", test.arr, got, test.expected)
    }
  }
}
