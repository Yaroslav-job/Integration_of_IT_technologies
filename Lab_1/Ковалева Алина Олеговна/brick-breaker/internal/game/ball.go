package game

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// Ball — структура, описывающая мяч.
type Ball struct {
	Image   *canvas.Image   // Графическое представление мяча
	dx      float64         // Смещение по X
	dy      float64         // Смещение по Y
	paddle  *Paddle         // Ссылка на платформу
	bricks  []*Brick        // Список кирпичей
	canvas  *fyne.Container // Канвас для отображения
	stopped bool            // Флаг остановки движения
}

// NewBall создает новый мяч и помещает его в центр окна
func NewBall(paddle *Paddle, bricks []*Brick, container *fyne.Container) *Ball {
	img := canvas.NewImageFromFile("internal/game/assets/ball.png")
	img.Resize(BallSize)
	img.Move(fyne.NewPos(
		(WindowSize.Width-BallSize.Width)/2,
		(WindowSize.Height-BallSize.Height)/2,
	))

	return &Ball{
		Image:  img,
		dx:     BallSpeed,
		dy:     -BallSpeed,
		paddle: paddle,
		bricks: bricks,
		canvas: container,
	}
}

// AddBall добавляет мяч в игру и запускает его движение в отдельной горутине
func AddBall(
	paddle *Paddle,
	bricks []*Brick,
	container *fyne.Container,
	onBallLost func(), // вызывается при потере мяча
	onCheckVictory func(), // вызывается при разрушении кирпича
) {
	ball := NewBall(paddle, bricks, container)

	activeBallsLock.Lock()
	activeBalls = append(activeBalls, ball)
	activeBallsLock.Unlock()

	container.Add(ball.Image)

	// Запускаем движение мяча в фоне
	go ball.Start(func() {
		// Удаляем мяч из активных
		activeBallsLock.Lock()
		for i, b := range activeBalls {
			if b == ball {
				activeBalls = append(activeBalls[:i], activeBalls[i+1:]...)
				break
			}
		}
		activeBallsLock.Unlock()

		container.Remove(ball.Image)
		onBallLost()
	}, onCheckVictory)
}

// Stop — останавливает мяч
func (b *Ball) Stop() {
	b.stopped = true
}

// Start — запускает логику движения мяча
func (b *Ball) Start(onBallLost func(), onCheckVictory func()) {
	for {
		if b.stopped {
			return
		}

		time.Sleep(16 * time.Millisecond) // ~60 FPS

		// Текущее и новое положение мяча
		pos := b.Image.Position()
		newX := pos.X + float32(b.dx)
		newY := pos.Y + float32(b.dy)

		// Отскок от стен
		if newX <= 0 || newX+BallSize.Width >= WindowSize.Width {
			b.dx *= -1
		}
		if newY <= 0 {
			b.dy *= -1
		}

		// Мяч упал вниз
		if newY+BallSize.Height >= WindowSize.Height {
			onBallLost()
			return
		}

		// Проверка столкновения с платформой
		paddlePos := b.paddle.Image.Position()
		paddleSize := b.paddle.Image.Size()
		ballSize := b.Image.Size()

		if newY+ballSize.Height >= paddlePos.Y &&
			newY+ballSize.Height <= paddlePos.Y+10 &&
			newX+ballSize.Width/2 >= paddlePos.X &&
			newX+ballSize.Width/2 <= paddlePos.X+paddleSize.Width {
			b.dy *= -1
			newY = paddlePos.Y - ballSize.Height
		}

		// Проверка столкновений с кирпичами
		for _, brick := range b.bricks {
			if !brick.Broken {
				pos := brick.Image.Position()
				size := brick.Image.Size()
				if newX+BallSize.Width >= pos.X && newX <= pos.X+size.Width &&
					newY+BallSize.Height >= pos.Y && newY <= pos.Y+size.Height {
					brick.Broken = true
					b.canvas.Remove(brick.Image)
					b.dy *= -1
					onCheckVictory()
					break
				}
			}
		}

		// Перемещаем мяч
		b.Image.Move(fyne.NewPos(newX, newY))
	}
}
