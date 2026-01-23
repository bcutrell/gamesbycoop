package sprites

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	spriteSize  = 15
	scaleFactor = 4
	outputSize  = spriteSize * scaleFactor
)

var smileyPattern = []string{
	".....XXXXX.....",
	"...XX#####XX...",
	"..X#########X..",
	".X###########X.",
	"X#OOOOOOOOOOO#X",
	"X#OOOO###OOOO#X",
	"X#############X",
	"X#############X",
	"X#############X",
	"X##########O##X",
	".X####O##OO##X.",
	"..X####OO###X..",
	"...XX#####XX...",
	".....XXXXX.....",
	"...............",
}

var charColors = map[rune]color.RGBA{
	'.': {0, 0, 0, 0},
	'X': {0, 0, 0, 255},
	'#': {255, 220, 0, 255},
	'O': {0, 0, 0, 255},
}

func NewSmileyFace() *ebiten.Image {
	smallImg := ebiten.NewImage(spriteSize, spriteSize)

	for y, row := range smileyPattern {
		for x, char := range row {
			if c, ok := charColors[char]; ok {
				smallImg.Set(x, y, c)
			}
		}
	}

	scaledImg := ebiten.NewImage(outputSize, outputSize)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(scaleFactor), float64(scaleFactor))
	op.Filter = ebiten.FilterNearest
	scaledImg.DrawImage(smallImg, op)

	return scaledImg
}
