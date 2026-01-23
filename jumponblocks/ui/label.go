package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Label struct {
	Text     string
	X, Y     float64
	Color    color.Color
	Centered bool
}

func (l *Label) Draw(screen *ebiten.Image, face text.Face) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(l.X, l.Y)
	op.ColorScale.ScaleWithColor(l.Color)

	if l.Centered {
		w, _ := text.Measure(l.Text, face, 0)
		op.GeoM.Translate(-w/2, 0)
	}

	text.Draw(screen, l.Text, face, op)
}
