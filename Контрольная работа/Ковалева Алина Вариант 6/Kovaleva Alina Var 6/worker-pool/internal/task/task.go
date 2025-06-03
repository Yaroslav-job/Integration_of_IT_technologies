// Вариант 6
// Package task определяет структуру и обработку задач для воркеров
// Разработчик: Ковалева Алина
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация

package task

import (
	"errors"
	"fmt"
	"net/url"
)

// Task представляет единицу работы для воркера
type Task struct {
	ID  int    // Уникальный идентификатор задачи
	URL string // URL для обработки
}

// Validate проверяет корректность URL в задаче
func (t *Task) Validate() error {
	if t.URL == "" {
		return errors.New("URL не может быть пустым")
	}
	
	_, err := url.ParseRequestURI(t.URL)
	if err != nil {
		return fmt.Errorf("некорректный URL: %w", err)
	}
	
	return nil
}

// Process имитирует обработку задачи
func (t *Task) Process() (string, error) {
	if err := t.Validate(); err != nil {
		return "", fmt.Errorf("ошибка валидации: %w", err)
	}
	
	// Имитация обработки
	result := fmt.Sprintf("Обработана задача %d: %s", t.ID, t.URL)
	return result, nil
}