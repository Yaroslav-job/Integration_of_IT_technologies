/*
Название:      network_test
Описание:      Unit-тесты для функций Sender и Receiver из пакета network.
Разработчик:   Code Copilot (https://chat.openai.com/gpts)
Лицензия:      GPLv3 — Свободное использование, модификация и распространение. Производные работы под GPLv3.

История изменений:
- 2025-04-03: Добавлены базовые и контролируемые тесты передачи и подтверждения сообщений.
*/

package network

import (
	"sync"
	"testing"
	"time"
)

// Тест Receiver с гарантированной доставкой (rand = 0.0)
func TestReceiver_AlwaysReceives(t *testing.T) {
	msgChan := make(chan Message, 3)
	ackChan := make(chan int, 3)

	// Предсказуемое поведение: не теряем сообщения
	originalRand := randFloat32
	randFloat32 = func() float32 { return 0.0 }
	defer func() { randFloat32 = originalRand }()

	msgs := []Message{
		{ID: 1, Body: "msg1"},
		{ID: 2, Body: "msg2"},
		{ID: 3, Body: "msg3"},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Receiver(msgChan, ackChan)
	}()

	for _, m := range msgs {
		msgChan <- m
	}
	close(msgChan)

	received := make(map[int]bool)
	timeout := time.After(1 * time.Second)

	for len(received) < len(msgs) {
		select {
		case ack := <-ackChan:
			received[ack] = true
		case <-timeout:
			t.Fatal("timeout waiting for ACKs")
		}
	}

	for _, m := range msgs {
		if !received[m.ID] {
			t.Errorf("message ID %d was not acknowledged", m.ID)
		}
	}

	wg.Wait()
}

// mockable rand function
var randFloat32 = func() float32 {
	return 0.5 // по умолчанию
}
