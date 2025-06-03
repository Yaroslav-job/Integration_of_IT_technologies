// Вариант 6
// Основной исполняемый файл для демонстрации пула воркеров
// Разработчик: Ковалева Алина
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация

package main

import (
	"fmt"
	"log"
	"worker-pool/internal/task"
	"worker-pool/internal/worker"
)

func main() {
	// Создаем пул из 3 воркеров
	pool := worker.New(3)
	pool.Start()

	// Горутина для обработки результатов
	go func() {
		for result := range pool.GetResults() {
			fmt.Println("Результат:", result)
		}
	}()

	// Создаем 10 задач
	tasks := []task.Task{
		{ID: 1, URL: "https://example.com/1"},
		{ID: 2, URL: "https://example.com/2"},
		{ID: 3, URL: "https://example.com/3"},
		{ID: 4, URL: "https://example.com/4"},
		{ID: 5, URL: "https://example.com/5"},
		{ID: 6, URL: "https://example.com/6"},
		{ID: 7, URL: "https://example.com/7"},
		{ID: 8, URL: "https://example.com/8"},
		{ID: 9, URL: "https://example.com/9"},
		{ID: 10, URL: "https://example.com/10"},
	}

	// Добавляем задачи в пул
	for _, t := range tasks {
		if err := t.Validate(); err != nil {
			log.Printf("Задача %d невалидна: %v", t.ID, err)
			continue
		}
		pool.AddTask(t)
	}

	// Ожидаем завершения работы
	pool.Stop()
	fmt.Println("Все задачи обработаны")
}
