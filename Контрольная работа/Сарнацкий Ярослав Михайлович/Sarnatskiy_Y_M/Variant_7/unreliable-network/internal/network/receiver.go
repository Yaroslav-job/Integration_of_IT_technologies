// receiver.go
// Название: receiver.go
// Описание: Логика приема сообщений с возможной потерей и отправкой подтверждения
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История: 03.04.2025 — создан файл

package network

import (
	"fmt"
	"math/rand"
	"time"
)

func Receiver(in <-chan Message, ack chan<- int, lossProbability float64) {
	rand.Seed(time.Now().UnixNano())
	for msg := range in {
		if rand.Float64() < lossProbability {
			fmt.Printf("[Receiver] Потеряно сообщение #%d\n", msg.ID)
			continue
		}
		fmt.Printf("[Receiver] Получено: %v\n", msg)
		ack <- msg.ID
	}
}