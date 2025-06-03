// Название: camelcase
// Краткое описание: Преобразование строки в верблюжий регистр (camelCase).
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений:
//    - 03.04.2025: Реализован алгоритм преобразования строки в camelCase.

package camelcase

import (
	"strings"
)

// ConvertToCamelCase преобразует строку в верблюжий регистр (camelCase).
func ConvertToCamelCase(input string) string {
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}

	// Начинаем с первого слова в нижнем регистре
	result := strings.ToLower(words[0])

	// Для каждого следующего слова первое буква заглавная
	for i := 1; i < len(words); i++ {
		word := words[i]
		// Преобразуем первую букву в заглавную, остальные в нижний
		result += strings.Title(strings.ToLower(word))
	}

	return result
}
