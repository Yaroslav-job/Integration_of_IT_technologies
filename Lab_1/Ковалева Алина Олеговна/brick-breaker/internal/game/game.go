package game

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowCenteredOverlayMenu(parent *fyne.Container, title string, onRestart func()) *fyne.Container {
	// Полупрозрачный фон
	overlayBg := canvas.NewRectangle(&color.NRGBA{R: 0, G: 0, B: 0, A: 180})
	overlayBg.Resize(WindowSize)

	// Меню с заголовком и кнопкой
	titleLabel := canvas.NewText(title, color.White)
	titleLabel.TextSize = 32
	titleLabel.Alignment = fyne.TextAlignCenter

	restartBtn := widget.NewButton("Restart", onRestart)

	menu := container.NewVBox(
		layout.NewSpacer(),
		titleLabel,
		widget.NewSeparator(),
		restartBtn,
		layout.NewSpacer(),
	)

	// Обернём в центрирующий контейнер
	centeredMenu := container.NewCenter(menu)

	// Финальный оверлей
	overlay := container.NewWithoutLayout(overlayBg, centeredMenu)
	centeredMenu.Resize(fyne.NewSize(300, 150))
	centeredMenu.Move(fyne.NewPos(
		(WindowSize.Width-300)/2,
		(WindowSize.Height-150)/2,
	))

	parent.Add(overlay)
	return overlay
}
