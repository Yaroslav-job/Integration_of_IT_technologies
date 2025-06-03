// Название: main
// Краткое описание: Основной файл для запуска программы.
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История изменений:
//    - 03.04.2025: Создан основной файл для запуска программы.

package main

import (
	"bufio"
	"fmt"
	"os"
	"camelcase-project/internal/camelcase"
)

func main() {
	// Ввод строки с консоли
	fmt.Println("Введите строку для преобразования:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Преобразуем строку в camelCase
	result := camelcase.ConvertToCamelCase(input)
	// Выводим результат
	fmt.Println("Converted String:", result)

	// Пример использования параллельного алгоритма
	parallelResult := camelcase.ConvertToCamelCaseParallel(input)
	fmt.Println("Parallel Converted String:", parallelResult)
}
