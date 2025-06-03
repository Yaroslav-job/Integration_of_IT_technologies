/*
Project: Контрольная работа. Алгоритмы и Go Routines.
Description: Находит наибольший общий префикс в заданном массиве строк.
Author: Никита Мякишев (вариант 3)
License: GPLv3
History:
  - [03.04.2025 17:02]: Initial task1
*/

package task1

// Функция для поиска наибольшего общего префикса
func LongestCommonPrefix(strs []string) string {
	// Если массив пустой, возвращаем пустую строку
	if len(strs) == 0 {
		return ""
	}

	// Берем первую строку как начальный префикс
	prefix := strs[0]

	// Проходим по всем оставшимся строкам в массиве
	for _, str := range strs[1:] {
		// Пока текущий префикс не является началом строки, уменьшаем его длину
		for len(prefix) > 0 && !startsWith(str, prefix) {
			prefix = prefix[:len(prefix)-1] // Срез последнего символа
		}
	}
	return prefix
}

// Вспомогательная функция для проверки, начинается ли строка с заданного префикса
func startsWith(str, prefix string) bool {
	return len(str) >= len(prefix) && str[:len(prefix)] == prefix
}
