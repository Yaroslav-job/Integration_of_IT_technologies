// main.go
// Название: main.go
// Описание: Точка входа, запускает горутины отправителя и получателя
// Разработчик: Сарнацкий Я.М.
// Лицензия: GPLv3
// История: 03.04.2025 — создан файл

package main

import (
    "testing"
    "time"
    "unreliable-network/internal/network"
)

func TestMainIntegration(t *testing.T) {
    msgChan := make(chan network.Message)
    ackChan := make(chan int)

    go network.Receiver(msgChan, ackChan, 0.0) // Без потерь
    go network.Sender(msgChan, ackChan)

    time.Sleep(2 * time.Second) // Даем время горутинам

    select {
    case ackID := <-ackChan:
        if ackID <= 0 {
            t.Errorf("Некорректное подтверждение: %d", ackID)
        }
    case <-time.After(3 * time.Second):
        t.Error("Подтверждение не получено")
    }
}