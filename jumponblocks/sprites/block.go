package sprites

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var ColorBlock = color.RGBA{139, 90, 43, 255} // Brown

func NewBlock(width, height int) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	img.Fill(ColorBlock)
	return img
}
