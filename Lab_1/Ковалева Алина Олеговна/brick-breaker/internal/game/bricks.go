package game

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// Brick — структура для представления кирпича
type Brick struct {
	Image  *canvas.Image
	Broken bool
}

// NewBricks создает массив кирпичей с разной текстурой
func NewBricks(mode string) []*Brick {

	var bricks []*Brick

	brickWidth := WindowSize.Width / float32(BrickColumns)
	brickHeight := float32(30)

	// Загрузка текстур
	textures := make([]string, 9)
	for i := 0; i < 9; i++ {
		textures[i] = fmt.Sprintf("internal/game/assets/brick%d.png", i+1)
	}

	lastUsed := make([]int, BrickColumns) // Для режима random

	for row := 0; row < BrickRows; row++ {
		for col := 0; col < BrickColumns; col++ {
			var texturePath string

			switch mode {
			case "diagonal":
				textureIndex := (row + col) % len(textures)
				texturePath = textures[textureIndex]

			case "random":
				var textureIndex int
				for {
					textureIndex = rand.Intn(len(textures))
					if col > 0 {
						prev := bricks[row*BrickColumns+col-1]
						if textures[textureIndex] == prev.Image.File {
							continue
						}
					}
					break
				}
				texturePath = textures[textureIndex]
				lastUsed[col] = textureIndex

			default:
				texturePath = "internal/game/assets/brick1.png"
			}

			img := canvas.NewImageFromFile(texturePath)
			img.Resize(fyne.NewSize(brickWidth, brickHeight))
			img.Move(fyne.NewPos(float32(col)*brickWidth, float32(row)*brickHeight))

			bricks = append(bricks, &Brick{Image: img})
		}
	}

	return bricks
}
