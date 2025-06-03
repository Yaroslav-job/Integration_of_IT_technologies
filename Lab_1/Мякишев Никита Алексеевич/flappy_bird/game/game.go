package game

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"golang.org/x/image/font/basicfont"

	"image/color"
)

var img_background *ebiten.Image

// Инициализация изображения
func init() {
	// Загрузить изображение с указанного пути
	var err error
	img_background, _, err = ebitenutil.NewImageFromFile("assets/bg.png")
	if err != nil {
		// Если произошла ошибка при загрузке, выводим ее и завершаем выполнение программы
		log.Fatal(err)
	}
}

type Game struct {
	bird  *Bird  // Птичка
	pipes *Pipes // Трубы
	timer *Timer // Таймер игры
	over  bool   // Флаг завершения игры
	score int    // Очки игрока
}

// NewGame создает новый объект Game.
// Возвращает указатель на созданный game.
func NewGame() *Game {
	return &Game{
		bird:  NewBird(),
		pipes: NewPipes(),
		timer: NewTimer(),
		over:  false,
	}
}

// Обновляет состояние игры, вызывается каждый кадр.
func (g *Game) Update() error {
	if g.over {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.Restart() // Рестарт игры при нажатии R
		}
		return nil
	}

	g.bird.Update()
	g.pipes.Update()

	// Проверяем столкновение с трубами
	for _, pipe := range g.pipes.pipes {
		if g.bird.CollidesWith(pipe) {
			g.over = true
		}
		if pipe.x+pipe.width < 100 && !pipe.isTop && !pipe.wasPassedPoint {
			pipe.wasPassedPoint = true
			// Если труба прошла точку, увеличиваем счёт
			g.score++
		}
	}

	// Проверяем, не упала ли птичка на землю
	if g.bird.y-g.bird.height >= 600 { // Земля находится на 600 пикселях, учитываем высоту птички
		g.over = true
	}

	return nil
}

// Отображает игровой экран
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(img_background, nil)

	g.bird.Draw(screen)
	g.pipes.Draw(screen)
	g.timer.Draw(screen)

	// Отображение очков
	scoreMessage := fmt.Sprintf("Score: %d", g.score)
	text.Draw(screen, scoreMessage, basicfont.Face7x13, 10, 20, color.White)

	if g.over {
		g.timer.Stop()

		// Создаем фон для текста (красный)
		textImage := ebiten.NewImage(225, 100)
		textImage.Fill(color.RGBA{255, 0, 0, 255})

		// Параметры для отображения фона
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(100, 250)
		screen.DrawImage(textImage, op)

		// Текст "Press R to Restart"
		message := "Press R to Restart"
		largeFont := basicfont.Face7x13
		// Рисуем текст поверх фона
		text.Draw(screen, message, largeFont, 150, 300, color.White) // Белый цвет для текста
	}
}

// Определяет размер игрового окна
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 600
}

// Метод для рестарта игры
func (g *Game) Restart() {
	// Сбрасываем позиции птички и труб
	g.bird = NewBird()
	g.pipes = NewPipes()
	g.timer = NewTimer()
	g.over = false
	g.score = 0
}
