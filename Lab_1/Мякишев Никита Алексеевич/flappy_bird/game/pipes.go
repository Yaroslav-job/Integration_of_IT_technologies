package game

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pipe struct {
	x, y           float64 // Координаты трубы
	width          float64 // Ширина трубы
	height         float64 // Высота трубы
	isTop          bool    // true - верхняя труба, false - нижняя труба
	wasPassedPoint bool    // Флаг, пройдена ли труба птичкой (для начисления очков)
}

type Pipes struct {
	pipes        []*Pipe       // Список труб
	lastPipeTime time.Time     // Время появления последней трубы
	pipeInterval time.Duration // Интервал между трубами
}

func NewPipes() *Pipes {
	rand.Seed(time.Now().UnixNano()) // инициализация генератора случайных чисел с использованием текущего времени, преобразованного в значение в наносекундах.
	pipes := &Pipes{}
	pipes.AddPipe()                      // Добавляем первую пару труб
	pipes.lastPipeTime = time.Now()      // Запоминаем время появления первой трубы
	pipes.pipeInterval = 2 * time.Second // Устанавливаем интервал появления труб (2 секунды)
	return pipes
}

// Обновляет позиции труб и добавляет новые.
func (p *Pipes) Update() {
	for _, pipe := range p.pipes {
		pipe.x -= 2 // Двигаем трубы влево
	}

	// Если первая труба ушла за экран, удаляем её и добавляем новую
	if len(p.pipes) > 0 && p.pipes[0].x < -50 {
		p.pipes = p.pipes[2:] // Удаляем первую пару труб
	}

	// Проверяем, прошло ли достаточно времени для добавления новой трубы
	if time.Since(p.lastPipeTime) > p.pipeInterval {
		p.AddPipe()                 // Добавляем новую трубу
		p.lastPipeTime = time.Now() // Обновляем время последнего добавления трубы
	}
}

// Отображает трубы на экране.
func (p *Pipes) Draw(screen *ebiten.Image) {
	for _, pipe := range p.pipes {
		pipeImg := ebiten.NewImage(int(pipe.width), int(pipe.height))
		pipeImg.Fill(color.RGBA{0, 255, 0, 255}) // Зеленые трубы
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pipe.x, pipe.y) // Устанавливаем позицию трубы
		screen.DrawImage(pipeImg, op)
	}
}

// Добавляем новую пару труб
func (p *Pipes) AddPipe() {
	gap := 150       // Разрыв между трубами
	minHeight := 50  // Минимальная высота трубы
	maxHeight := 250 // Максимальная высота трубы

	// Случайное смещение трубы (от -50 до +50 пикселей)
	offset := rand.Intn(101) - 50

	// Генерация высоты верхней трубы с учетом смещения
	topHeight := rand.Intn(maxHeight-minHeight) + minHeight + offset
	if topHeight < minHeight {
		topHeight = minHeight // Обеспечиваем, чтобы труба не опустилась слишком низко
	} else if topHeight > maxHeight {
		topHeight = maxHeight // Обеспечиваем, чтобы труба не поднялась слишком высоко
	}

	// Расчет нижней трубы
	bottomY := float64(topHeight + gap)
	bottomHeight := 600 - bottomY

	// Добавляем верхнюю и нижнюю трубу в массив
	p.pipes = append(p.pipes,
		&Pipe{x: 400, y: 0, width: 50, height: float64(topHeight), isTop: true},
		&Pipe{x: 400, y: bottomY, width: 50, height: float64(bottomHeight), isTop: false},
	)

}
