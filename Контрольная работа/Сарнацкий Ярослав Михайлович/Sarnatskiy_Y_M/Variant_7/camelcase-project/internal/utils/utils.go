// Название: utils
// Краткое описание: Вспомогательные функции для обработки данных.
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений:
//    - 03.04.2025: Реализованы вспомогательные функции.

package utils

import "strings"

// ToLowerCase принимает строку и возвращает ее в нижнем регистре.
func ToLowerCase(input string) string {
	return strings.ToLower(input)
}
