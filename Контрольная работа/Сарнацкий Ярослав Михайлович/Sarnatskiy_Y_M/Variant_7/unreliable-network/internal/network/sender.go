// sender.go
// Название: sender.go
// Описание: Логика отправки сообщений с подтверждением и повторами
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История: 03.04.2025 — создан файл

package network

import (
	"fmt"
	"time"
)

func Sender(out chan<- Message, ack <-chan int) {
	messageID := 1
	for {
		msg := Message{
			ID:      messageID,
			Body:    fmt.Sprintf("Сообщение #%d", messageID),
			Created: time.Now(),
		}
		out <- msg
		fmt.Printf("[Sender] Отправлено: %v\n", msg)

		select {
		case ackID := <-ack:
			if ackID == messageID {
				fmt.Printf("[Sender] Получено подтверждение для сообщения #%d\n", ackID)
				messageID++
				continue
			}
		case <-time.After(1 * time.Second):
			fmt.Printf("[Sender] Повторная отправка сообщения #%d\n", messageID)
		}
	}
}