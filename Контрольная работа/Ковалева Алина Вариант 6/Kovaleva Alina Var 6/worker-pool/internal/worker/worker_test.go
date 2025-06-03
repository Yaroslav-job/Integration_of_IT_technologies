// Вариант 6
// Тесты для пакета worker
// Разработчик: Ковалева Алина
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация

package worker

import (
	"fmt"
	"sync"
	"testing"
	"worker-pool/internal/task"
)

func TestWorkerPool(t *testing.T) {
	pool := New(2)
	pool.Start()

	var wg sync.WaitGroup
	wg.Add(1)

	// Горутина для чтения результатов
	go func() {
		defer wg.Done()
		for range pool.GetResults() {
			// Тест только проверяет что система не падает
		}
	}()

	// Добавляем тестовые задачи
	for i := 0; i < 5; i++ {
		pool.AddTask(task.Task{
			ID:  i + 1,
			URL: fmt.Sprintf("https://test%d.com", i+1),
		})
	}

	pool.Stop()
	wg.Wait()
}