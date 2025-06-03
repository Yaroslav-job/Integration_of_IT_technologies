package main

import (
	"brick-breaker/internal/game"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	win := a.NewWindow("Brick Breaker")

	game.StartGame(win)                 // сначала устанавливаем контент
	win.Resize(game.ExternalWindowSize) // потом — точный размер
	win.SetFixedSize(true)              // запрещаем изменение размера

	win.ShowAndRun()
}
