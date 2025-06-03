// Вариант 6
// Package worker реализует пул горутин для обработки задач
// Разработчик: Ковалева Алина
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация

package worker

import (
	"log"
	"sync"
	"worker-pool/internal/task"
)

// WorkerPool управляет пулом воркеров
type WorkerPool struct {
	taskChan    chan task.Task
	resultChan  chan string
	workerCount int
	wg          sync.WaitGroup
}

// New создает новый пул воркеров
func New(workerCount int) *WorkerPool {
	return &WorkerPool{
		taskChan:    make(chan task.Task, 100),
		resultChan:  make(chan string, 100),
		workerCount: workerCount,
	}
}

// Start запускает воркеров
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i + 1)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	
	for t := range wp.taskChan {
		result, err := t.Process()
		if err != nil {
			log.Printf("Воркер %d: ошибка обработки задачи %d: %v", id, t.ID, err)
			continue
		}
		wp.resultChan <- result
	}
}

// AddTask добавляет задачу в пул
func (wp *WorkerPool) AddTask(t task.Task) {
	wp.taskChan <- t
}

// GetResults возвращает канал с результатами
func (wp *WorkerPool) GetResults() <-chan string {
	return wp.resultChan
}

// Stop останавливает пул
func (wp *WorkerPool) Stop() {
	close(wp.taskChan)
	wp.wg.Wait()
	close(wp.resultChan)
}