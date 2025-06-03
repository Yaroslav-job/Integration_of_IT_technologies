/*
Название:      network_main
Описание:      Пример использования пакета network — надёжная доставка сообщений по ненадёжному каналу.
Разработчик:   Колычев Никита
Лицензия:      GPLv3 — Свободное использование, модификация и распространение. Производные работы под GPLv3.

История изменений:
- 2025-04-03: Демонстрация передачи сообщений с подтверждениями.
*/

package main

import (
	"fmt"
	"golangassignment/modules/network"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	msgChan := make(chan network.Message)
	ackChan := make(chan int)

	messages := []network.Message{
		{ID: 1, Body: "Hello"},
		{ID: 2, Body: "World"},
		{ID: 3, Body: "from"},
		{ID: 4, Body: "Code"},
		{ID: 5, Body: "Copilot"},
	}

	go network.Receiver(msgChan, ackChan)
	network.Sender(messages, msgChan, ackChan)

	fmt.Println("✅ Все сообщения подтверждены")
}
