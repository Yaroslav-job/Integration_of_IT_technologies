package game

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// Paddle — структура платформы
type Paddle struct {
	Image *canvas.Image
	Width float32
}

// NewPaddle создает платформу и размещает ее у нижнего края
func NewPaddle() *Paddle {
	img := canvas.NewImageFromFile("internal/game/assets/paddle.png")
	img.Resize(PaddleSize)
	img.Move(fyne.NewPos(
		(WindowSize.Width-PaddleSize.Width)/2,
		WindowSize.Height-PaddleSize.Height,
	))

	return &Paddle{Image: img, Width: PaddleSize.Width}
}

// MoveLeft перемещает платформу влево
func (p *Paddle) MoveLeft() {
	pos := p.Image.Position()
	newX := pos.X - 20
	if newX < 0 {
		newX = 0
	}
	p.Image.Move(fyne.NewPos(newX, pos.Y))
	canvas.Refresh(p.Image)
}

// MoveRight перемещает платформу вправо
func (p *Paddle) MoveRight() {
	pos := p.Image.Position()
	newX := pos.X + 20
	if newX+p.Width > WindowSize.Width {
		newX = WindowSize.Width - p.Width
	}
	p.Image.Move(fyne.NewPos(newX, pos.Y))
	canvas.Refresh(p.Image)
}
