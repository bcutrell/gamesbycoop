package ui

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	Text       string
	X, Y       float64
	Width      float64
	Height     float64
	Color      color.Color
	HoverColor color.Color
	TextColor  color.Color
	hovered    bool
	OnClick    func()
}

func NewButton(txt string, x, y, w, h float64, clr, hoverClr, txtClr color.Color) *Button {
	return &Button{
		Text:       txt,
		X:          x,
		Y:          y,
		Width:      w,
		Height:     h,
		Color:      clr,
		HoverColor: hoverClr,
		TextColor:  txtClr,
	}
}

func (b *Button) Update() {
	mx, my := ebiten.CursorPosition()
	rect := image.Rect(int(b.X), int(b.Y), int(b.X+b.Width), int(b.Y+b.Height))
	b.hovered = image.Pt(mx, my).In(rect)

	if b.hovered && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if b.OnClick != nil {
			b.OnClick()
		}
	}
}

func (b *Button) Draw(screen *ebiten.Image, face text.Face) {
	clr := b.Color
	if b.hovered {
		clr = b.HoverColor
	}

	r, g, bl, a := clr.RGBA()
	vector.FillRect(screen, float32(b.X), float32(b.Y), float32(b.Width), float32(b.Height),
		color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(bl >> 8), uint8(a >> 8)}, false)

	op := &text.DrawOptions{}
	tw, th := text.Measure(b.Text, face, 0)
	op.GeoM.Translate(b.X+b.Width/2-tw/2, b.Y+b.Height/2-th/2)
	op.ColorScale.ScaleWithColor(b.TextColor)
	text.Draw(screen, b.Text, face, op)
}
