// network_test.go
// Название: network_test.go
// Описание: Модульные тесты для отправки и получения сообщений
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История: 03.04.2025 — создан файл

package network

import (
	"testing"
	"time"
)

func TestMessageStruct(t *testing.T) {
	msg := Message{ID: 1, Body: "test", Created: time.Now()}
	if msg.ID != 1 || msg.Body != "test" {
		t.Errorf("Message структура работает некорректно: %+v", msg)
	}
}

func TestAckSimulation(t *testing.T) {
	msgChan := make(chan Message, 1)
	ackChan := make(chan int, 1)

	go Receiver(msgChan, ackChan, 0.0) // 0% потеря
	msg := Message{ID: 42, Body: "ping", Created: time.Now()}
	msgChan <- msg

	select {
	case ackID := <-ackChan:
		if ackID != 42 {
			t.Errorf("Ожидалось подтверждение 42, получено %d", ackID)
		}
	case <-time.After(2 * time.Second):
		t.Error("Подтверждение не получено")
	}
}
