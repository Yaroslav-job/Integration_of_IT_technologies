// Package generator генерирует случайные строки и отправляет их в канал.
package generator

import (
	"math/rand"
	"time"
)

// Константы
const (
	Charset    = "abcdefghijklmnopqrstuvwxyz"
	StrLen     = 5
	NumStrings = 10
)

// GenerateStrings создает случайные строки и отправляет их в канал.
func GenerateStrings(ch chan<- string) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < NumStrings; i++ {
		ch <- randomString()
	}
	close(ch)
}

// randomString генерирует случайную строку длиной StrLen.
func randomString() string {
	b := make([]byte, StrLen)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}
	return string(b)
}
