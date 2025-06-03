/*
Package camelcase

Название:      camelcase
Описание:      Преобразование строк в camelCase без использования регулярных выражений.
Разработчик:   Колычев Никита
Лицензия:      GPLv3 — Свободное использование, модификация и распространение. Любые производные работы должны оставаться под GPLv3.

История изменений:
- 2025-04-03: Первая реализация с поддержкой ошибок и модульной структурой.
*/

package camelcase

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

/*
// Convert преобразует строку в camelCase.
// Возвращает ошибку, если входная строка пустая или содержит недопустимые символы.
*/
func Convert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("input string is empty or only spaces")
	}

	hasLetter := false
	for _, char := range trimmed {
		if unicode.IsLetter(char) {
			hasLetter = true
			break
		}
	}
	if !hasLetter {
		return "", errors.New("input must contain at least one letter")
	}

	var result strings.Builder
	capitalizeNext := false
	firstWord := true

	for _, char := range trimmed {
		if unicode.IsSpace(char) {
			capitalizeNext = true
			continue
		}

		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return "", fmt.Errorf("invalid character in input: '%c'", char)
		}

		if firstWord {
			result.WriteRune(unicode.ToLower(char))
			firstWord = false
			continue
		}

		if capitalizeNext {
			result.WriteRune(unicode.ToUpper(char))
			capitalizeNext = false
		} else {
			result.WriteRune(unicode.ToLower(char))
		}
	}

	return result.String(), nil
}
