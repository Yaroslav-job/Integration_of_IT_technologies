// arrayRotation_test.go
// Модульные тесты для функции циклического сдвига массива
// Разработчик: Nikita Pashin
// Лицензия: GPLv3
// История изменений:
// 2025-04-03: Первая версия

package rotation

import (
	"reflect"
	"testing"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		n        int
		expected []int
	}{
		{
			name:     "Пример из задания",
			arr:      []int{1, 2, 3, 4, 5},
			n:        2,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "Пустой массив",
			arr:      []int{},
			n:        3,
			expected: []int{},
		},
		{
			name:     "Массив из одного элемента",
			arr:      []int{42},
			n:        5,
			expected: []int{42},
		},
		{
			name:     "Сдвиг на 0 позиций",
			arr:      []int{1, 2, 3},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Сдвиг на кратное длине число",
			arr:      []int{1, 2, 3, 4},
			n:        4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Сдвиг на число больше длины",
			arr:      []int{1, 2, 3, 4},
			n:        6,
			expected: []int{3, 4, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RotateRight(tt.arr, tt.n)
			if !reflect.DeepEqual(tt.arr, tt.expected) {
				t.Errorf("RotateRight() = %v, ожидаем %v", tt.arr, tt.expected)
			}
		})
	}
}