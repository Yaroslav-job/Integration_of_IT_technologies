package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var img *ebiten.Image

// Инициализация изображения
func init() {
	// Загрузить изображение с указанного пути
	var err error
	img, _, err = ebitenutil.NewImageFromFile("assets/bird.png")
	if err != nil {
		// Если произошла ошибка при загрузке, выводим ее и завершаем выполнение программы
		log.Fatal(err)
	}
}

type Bird struct {
	x, y   float64 // Координаты птички
	vy     float64 // Вертикальная скорость
	width  float64 // Ширина птички
	height float64 // Высота птички
}

// Создает новую птичку в начальном положении
func NewBird() *Bird {
	return &Bird{x: 100, y: 200, width: 30, height: 30}
}

// Обновляет положение птички в каждом кадре игры.
func (b *Bird) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		b.vy = -4 // Прыжок вверх при нажатии пробела
	}
	b.y += b.vy // Двигаем птичку по вертикали
	b.vy += 0.2 // Применяем гравитацию
}

// Функция Draw рисует птичку на экране
func (b *Bird) Draw(screen *ebiten.Image) {
	// Создаем объект опций для настройки рисования изображения
	op := &ebiten.DrawImageOptions{}

	// Масштабируем изображение до размера птицы с учетом небольшого увеличения на 20 пикселей
	// Мы берем ширину и высоту изображения птицы (b.width и b.height), увеличиваем их на 20 и затем вычисляем коэффициент масштабирования
	op.GeoM.Scale(float64(b.width+20)/float64(img.Bounds().Dx()), float64(b.height+20)/float64(img.Bounds().Dy()))

	// Устанавливаем позицию, где будет нарисована птичка
	// b.x и b.y — это координаты верхнего левого угла, куда мы хотим отобразить птичку
	op.GeoM.Translate(b.x, b.y) // Перемещаем изображение в точку (b.x, b.y)

	// Отрисовываем изображение птицы с применением трансформации (масштабирования и перевода)
	screen.DrawImage(img, op)
}

// Функция проверки столкновения с трубой
func (b *Bird) CollidesWith(pipe *Pipe) bool {
	// Границы птички
	birdLeft := b.x
	birdRight := b.x + b.width
	birdTop := b.y
	birdBottom := b.y + b.height

	// Границы трубы
	pipeLeft := pipe.x
	pipeRight := pipe.x + pipe.width
	pipeTop := pipe.y
	pipeBottom := pipe.y + pipe.height

	// Проверяем столкновение с верхней трубой
	if pipe.isTop {
		if birdRight > pipeLeft && birdLeft < pipeRight && birdTop < pipeBottom {
			return true
		}
	} else {
		// Проверяем столкновение с нижней трубой
		if birdRight > pipeLeft && birdLeft < pipeRight && birdBottom > pipeTop {
			return true
		}
	}
	return false
}
