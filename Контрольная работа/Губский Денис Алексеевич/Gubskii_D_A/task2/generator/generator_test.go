// generator/generator_test.go
// Тесты для пакета генератора случайных чисел
//
// Разработчик: Губский Денис Алексеевич
// Лицензия: GPLv3
// История изменений:
//  - v 0.1: Создание проекта, реализация задачи с использованием горутин и каналов
//  - v 0.2: Добавление тестов и файлов Makefile и LICENSE.md

package generator

import (
	"testing" // Импортируем пакет для написания тестов
)

func TestGenerateRandomNumbers(t *testing.T) {
	ch := make(chan int) // Создаем канал для передачи чисел
	count := 20          // Указываем количество чисел для генерации

	// Запускаем генерацию случайных чисел в отдельной горутине
	go GenerateRandomNumbers(ch, count)

	// Проверяем, что канал содержит 20 чисел
	numCount := 0
	for num := range ch {
		numCount++ // Увеличиваем счетчик чисел
		if num < 1 || num > 100 {
			t.Errorf("Сгенерированное число %d не в пределах от 1 до 100", num)
		}
	}

	// Проверяем, что количество чисел соответствует ожидаемому
	if numCount != count {
		t.Errorf("Ожидалось %d чисел, а получено %d", count, numCount)
	}

	// Тест с count = 0
	chZero := make(chan int) // Создаем новый канал
	go GenerateRandomNumbers(chZero, 0)
	if _, ok := <-chZero; ok {
		t.Error("Канал должен быть закрыт при count = 0")
	}

	// Тест с большим count
	chLarge := make(chan int) // Создаем новый канал
	go GenerateRandomNumbers(chLarge, 1000)
	numCount = 0
	for range chLarge {
		numCount++ // Увеличиваем счетчик чисел
	}
	if numCount != 1000 {
		t.Errorf("Ожидалось 1000 чисел, а получено %d", numCount)
	}
}
