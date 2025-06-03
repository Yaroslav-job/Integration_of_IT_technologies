package game

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// StartGame инициализирует и запускает новую игру
func StartGame(win fyne.Window) {
	// Создаем черный фон
	bg := canvas.NewRectangle(&color.NRGBA{R: 0, G: 0, B: 0, A: 255})
	bg.Resize(WindowSize)
	bg.StrokeColor = color.Black
	bg.StrokeWidth = 2

	// Основной контейнер для всех игровых объектов
	gameCanvas := container.NewWithoutLayout()
	gameCanvas.Resize(WindowSize)
	gameCanvas.Add(bg)

	// Создание платформы игрока
	paddle := NewPaddle()
	gameCanvas.Add(paddle.Image)

	// Создание кирпичей
	bricks := NewBricks("diagonal")
	for _, brick := range bricks {
		gameCanvas.Add(brick.Image)
	}

	// Инициализация списка мячей и флага окончания игры
	activeBalls = []*Ball{}
	gameEnded := false
	var currentOverlay *fyne.Container

	// Показать меню окончания игры (победа/проигрыш)
	showMenu := func(title string) {
		gameEnded = true

		// Остановить и удалить все мячи
		activeBallsLock.Lock()
		for _, b := range activeBalls {
			b.Stop()
			gameCanvas.Remove(b.Image)
		}
		activeBalls = nil
		activeBallsLock.Unlock()

		// Удалить старое меню, если есть
		if currentOverlay != nil {
			gameCanvas.Remove(currentOverlay)
		}

		// Показать новое меню с заголовком и кнопкой перезапуска
		currentOverlay = ShowCenteredOverlayMenu(gameCanvas, title, func() {
			StartGame(win)
		})
	}

	// Проверка победы (все кирпичи разрушены)
	checkVictory := func() {
		if gameEnded {
			return
		}
		allBroken := true
		for _, brick := range bricks {
			if !brick.Broken {
				allBroken = false
				break
			}
		}
		if allBroken {
			showMenu("Victory!")
		}
	}

	// Проверка проигрыша (все мячи потеряны)
	checkGameOver := func() {
		if gameEnded {
			return
		}
		activeBallsLock.Lock()
		count := len(activeBalls)
		activeBallsLock.Unlock()

		if count == 0 {
			showMenu("Game Over")
		}
	}

	// Добавляем первый мяч в игру
	AddBall(paddle, bricks, gameCanvas, checkGameOver, checkVictory)

	// Горутину для усложнения игры со временем:
	// каждые 15 секунд — ускоряем мячи, уменьшаем платформу, добавляем мяч
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			if gameEnded {
				return
			}

			// Увеличиваем скорость всех мячей на 10%
			activeBallsLock.Lock()
			for _, b := range activeBalls {
				b.dx *= 1.2
				b.dy *= 1.2
			}
			activeBallsLock.Unlock()

			// Уменьшаем ширину платформы на 10% (минимум 40 пикселей)
			currentSize := paddle.Image.Size()
			newWidth := currentSize.Width * 0.9
			if newWidth < 40 {
				newWidth = 40
			}
			paddle.Width = newWidth
			paddle.Image.Resize(fyne.NewSize(newWidth, currentSize.Height))

			// Добавляем дополнительный мяч (если еще не достигнут лимит)
			activeBallsLock.Lock()
			count := len(activeBalls)
			activeBallsLock.Unlock()

			if count < maxBalls {
				AddBall(paddle, bricks, gameCanvas, checkGameOver, checkVictory)
			}
		}
	}()

	// Отображаем канвас в окне
	win.SetContent(gameCanvas)

	// Обработка клавиш (влево/вправо)
	win.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyLeft:
			paddle.MoveLeft()
		case fyne.KeyRight:
			paddle.MoveRight()
		}
	})
}
