package game

import (
	"fmt"
	"image/color"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

// Timer управляет отсчетом времени в игре.
// Он работает в отдельной горутине и обновляет игровое время каждую секунду.
type Timer struct {
	startTime time.Time     // Время начала игры
	gameTime  int           // Время игры в секундах
	mutex     sync.Mutex    // Мьютекс для потокобезопасного обновления времени
	stopChan  chan struct{} // Канал для остановки таймера
	once      sync.Once     // Объект для гарантии однократного закрытия канала
}

// runGameTimer запускает горутину, которая обновляет время игры каждую секунду.
// Таймер работает до тех пор, пока не будет получен сигнал из stopChan.
func (t *Timer) runGameTimer() {
	ticker := time.NewTicker(1 * time.Second) // создание тикера, который срабатывает каждую секунду
	defer ticker.Stop()                       // остановка тикера при заврешениии горутины

	for {
		select {
		case <-ticker.C:
			t.mutex.Lock()
			t.gameTime = int(time.Since(t.startTime).Seconds()) // Обновляем время игры
			t.mutex.Unlock()

		case <-t.stopChan:
			return // Завершаем горутину
		}
	}
}

// Draw отображает текущее время игры на экране.
// Выводится текст "Time: X sec" в левом верхнем углу экрана.
func (t *Timer) Draw(screen *ebiten.Image) {
	t.mutex.Lock()
	gameTime := t.gameTime
	t.mutex.Unlock()

	// Отображаем время игры белым текстом
	text.Draw(screen, fmt.Sprintf("Time: %d sec", gameTime), basicfont.Face7x13, 10, 40, color.White)

}

// Stop останавливает таймер, гарантируя, что stopChan будет закрыт только один раз.
func (t *Timer) Stop() {
	t.once.Do(func() {
		close(t.stopChan)
	})
}

// NewTimer создает новый объект Timer и запускает его горутину.
// Возвращает указатель на созданный таймер.
func NewTimer() *Timer {
	timer := &Timer{
		startTime: time.Now(),
		stopChan:  make(chan struct{}),
	}
	go timer.runGameTimer()
	return timer
}
