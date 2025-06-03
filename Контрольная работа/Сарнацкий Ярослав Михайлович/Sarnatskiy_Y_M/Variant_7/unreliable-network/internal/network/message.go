// message.go
// Название: message.go
// Описание: Определение структуры сообщения
// Разработчик: Саранцкий Я.М.
// Лицензия: GPLv3
// История: 03.04.2025 — создан файл

package network

import "time"

type Message struct {
	ID      int
	Body    string
	Created time.Time
}