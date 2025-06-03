// Вариант 6
// Тесты для пакета task
// Разработчик: Ковалева Алина
// Лицензия: GPLv3
// История изменений:
// - 03.04.2025: Первоначальная реализация

package task

import (
	"testing"
)

func TestTaskValidation(t *testing.T) {
	tests := []struct {
		name    string
		task    Task
		wantErr bool
	}{
		{
			name:    "Корректный URL",
			task:    Task{ID: 1, URL: "https://example.com"},
			wantErr: false,
		},
		{
			name:    "Пустой URL",
			task:    Task{ID: 2, URL: ""},
			wantErr: true,
		},
		{
			name:    "Некорректный URL",
			task:    Task{ID: 3, URL: "invalid-url"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.task.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTaskProcessing(t *testing.T) {
	validTask := Task{ID: 1, URL: "https://valid.com"}
	result, err := validTask.Process()
	if err != nil {
		t.Fatalf("Process() неожиданная ошибка: %v", err)
	}
	if result == "" {
		t.Error("Process() вернул пустой результат")
	}
}