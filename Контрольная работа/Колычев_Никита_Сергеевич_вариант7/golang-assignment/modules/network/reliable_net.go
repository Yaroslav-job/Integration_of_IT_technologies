/*
Package network

Название:      network
Описание:      Имитация ненадежной сети с подтверждением доставки сообщений.
Разработчик:   Колычев Никита
Лицензия:      GPLv3 — Свободное использование, модификация и распространение. Производные работы под GPLv3.

История изменений:
- 2025-04-03: Первая версия с подтверждением сообщений и вероятностью потери.
*/

package network

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	ID   int
	Body string
}

// Sender отправляет сообщения и повторяет при отсутствии подтверждения.
func Sender(messages []Message, msgChan chan<- Message, ackChan <-chan int) {
	for _, msg := range messages {
		for {
			fmt.Printf("Отправка сообщения #%d: %s\n", msg.ID, msg.Body)
			msgChan <- msg

			select {
			case ackID := <-ackChan:
				if ackID == msg.ID {
					fmt.Printf("Подтверждение получено для #%d\n", msg.ID)
					time.Sleep(100 * time.Millisecond)
					break
				}
			case <-time.After(500 * time.Millisecond):
				fmt.Printf("Таймаут для #%d, повторная отправка...\n", msg.ID)
				continue
			}
			break
		}
	}
	close(msgChan)
}

// Receiver принимает сообщения с вероятностью потери и отправляет подтверждение.
func Receiver(msgChan <-chan Message, ackChan chan<- int) {
	for msg := range msgChan {
		if rand.Float32() < 0.2 {
			fmt.Printf("Потеряно сообщение #%d: %s\n", msg.ID, msg.Body)
			continue
		}
		fmt.Printf("Получено сообщение #%d: %s\n", msg.ID, msg.Body)
		time.Sleep(100 * time.Millisecond)
		ackChan <- msg.ID
	}
}
