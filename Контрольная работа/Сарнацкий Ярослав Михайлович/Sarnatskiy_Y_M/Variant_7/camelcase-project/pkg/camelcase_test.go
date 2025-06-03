// Название: camelcase_test
// Краткое описание: Модульные тесты для преобразования строки в camelCase.
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений:
//    - 03.04.2025: Добавлены тесты для проверки корректности работы алгоритма.

package camelcase_test

import (
	"camelcase-project/internal/camelcase" // Путь к пакету, где определены функции
	"testing"
)

// Тест для базового случая
func TestConvertToCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world example", "helloWorldExample"},
		{"Go programming language", "goProgrammingLanguage"},
		{"HELLO WORLD", "helloWorld"},
		{"singleword", "singleword"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := camelcase.ConvertToCamelCase(test.input)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Тест для параллельной версии
func TestConvertToCamelCaseParallel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world example", "helloWorldExample"},
		{"Go programming language", "goProgrammingLanguage"},
		{"HELLO WORLD", "helloWorld"},
		{"singleword", "singleword"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := camelcase.ConvertToCamelCaseParallel(test.input)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}
