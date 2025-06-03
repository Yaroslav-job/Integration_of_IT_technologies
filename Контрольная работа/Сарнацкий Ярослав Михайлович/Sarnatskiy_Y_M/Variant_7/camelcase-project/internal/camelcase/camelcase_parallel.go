// Название: camelcase_parallel
// Краткое описание: Параллельное преобразование строки в верблюжий регистр (camelCase) с использованием Go-рутину.
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений:
//    - 03.04.2025: Добавлена параллельная обработка с использованием Go-рутину.

package camelcase

import (
	"strings"
	"sync"
)

// ConvertToCamelCaseParallel параллельно преобразует строку в верблюжий регистр (camelCase).
func ConvertToCamelCaseParallel(input string) string {
	words := strings.Fields(input)
	if len(words) == 0 {
		return ""
	}

	var wg sync.WaitGroup
	results := make([]string, len(words))

	// Обрабатываем первое слово в главной рутине
	results[0] = strings.ToLower(words[0])
	wg.Add(len(words) - 1)

	// Обрабатываем остальные слова параллельно
	for i := 1; i < len(words); i++ {
		go func(index int, word string) {
			defer wg.Done()
			results[index] = strings.Title(strings.ToLower(word))
		}(i, words[i])
	}

	// Ждем завершения всех рутин
	wg.Wait()

	// Собираем результат
	return strings.Join(results, "")
}
