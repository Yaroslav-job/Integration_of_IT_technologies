/*
Название:      camelcase_test
Описание:      Unit-тесты для функции Convert из пакета camelcase.
Разработчик:   Колычев Никита
Лицензия:      GPLv3 — Свободное использование, модификация и распространение. Любые производные работы должны оставаться под GPLv3.

История изменений:
- 2025-04-03: Добавлены базовые и граничные тестовые случаи.
*/

package camelcase

import (
	"testing"
)

func TestConvert_ValidInputs(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"   spaced out  ", "spacedOut"},
		{"multiple words in sentence", "multipleWordsInSentence"},
		{"Camel CASE Test", "camelCaseTest"},
		{"one", "one"},
		{"with123 numbers", "with123Numbers"},
	}

	for _, tt := range tests {
		result, err := Convert(tt.input)
		if err != nil {
			t.Errorf("unexpected error for input %q: %v", tt.input, err)
			continue
		}
		if result != tt.expected {
			t.Errorf("Convert(%q) = %q, expected %q", tt.input, result, tt.expected)
		}
	}
}

func TestConvert_InvalidInputs(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"", true},
		{"     ", true},
		{"!!!", true},
		{"@@hello", true},
		{"   123 !@#   ", true},
	}

	for _, tt := range tests {
		_, err := Convert(tt.input)
		if (err != nil) != tt.expectError {
			t.Errorf("Convert(%q) error = %v, expectError = %v", tt.input, err, tt.expectError)
		}
	}
}
