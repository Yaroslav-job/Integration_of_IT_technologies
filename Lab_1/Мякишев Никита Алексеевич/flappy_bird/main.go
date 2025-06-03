package main

import (
	"log"

	"flappy_bird/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()                  // Создаем новый экземпляр игры
	ebiten.SetWindowSize(400, 600)       // Устанавливаем размер окна 400x600 пикселей
	ebiten.SetWindowTitle("Flappy Bird") // Устанавливаем заголовок окна

	if err := ebiten.RunGame(g); err != nil { // Запускаем игру
		log.Fatal(err) // Выводим ошибку и завершаем программу, если запуск не удался
	}
}
