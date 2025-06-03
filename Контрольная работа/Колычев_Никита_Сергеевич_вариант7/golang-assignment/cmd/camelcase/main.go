/*
Название:      camelcase_main
Описание:      Демонстрация использования модуля camelcase для преобразования строк в camelCase.
Разработчик:   Колычев Никита
Лицензия:      GPLv3 — Свободное использование, модификация и распространение. Любые производные работы должны оставаться под GPLv3.

История изменений:
- 2025-04-03: Базовая демонстрационная программа с примером использования функции Convert из модуля camelcase.
*/

package main

import (
	"fmt"
	"golangassignment/modules/camelcase"
)

func main() {
	examples := []string{
		"hello world example",
		" spaced  input ",
		"",
		"   ",
		"42 test case",
		"hello @world",
	}

	for _, ex := range examples {
		result, err := camelcase.Convert(ex)
		if err != nil {
			fmt.Printf("Input: %-20q ➜ Error: %v\n", ex, err)
		} else {
			fmt.Printf("Input: %-20q ➜ CamelCase: %s\n", ex, result)
		}
	}
}
