package screens

import (
	"image/color"

	"jumponblocks/sprites"
	"jumponblocks/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const (
	ScreenWidth  = 400
	ScreenHeight = 600
)

var (
	ColorText        = color.RGBA{255, 255, 255, 255}
	ColorButton      = color.RGBA{76, 175, 80, 255}
	ColorButtonHover = color.RGBA{102, 187, 106, 255}
)

type CoverScreen struct {
	levelLabel     *ui.Label
	playButton     *ui.Button
	bestScoreLabel *ui.Label
	pickCharLabel  *ui.Label
	incomeLabel    *ui.Label
	character      *ebiten.Image
}

func NewCoverScreen(onPlay func()) *CoverScreen {
	centerX := float64(ScreenWidth) / 2

	cs := &CoverScreen{
		levelLabel: &ui.Label{
			Text:     "Level 1",
			X:        centerX,
			Y:        60,
			Color:    ColorText,
			Centered: true,
		},
		playButton: ui.NewButton(
			"PLAY",
			centerX-60, 100,
			120, 50,
			ColorButton,
			ColorButtonHover,
			ColorText,
		),
		bestScoreLabel: &ui.Label{
			Text:  "Best: 0",
			X:     20,
			Y:     ScreenHeight - 80,
			Color: ColorText,
		},
		pickCharLabel: &ui.Label{
			Text:  "Player",
			X:     ScreenWidth - 95,
			Y:     ScreenHeight - 80,
			Color: ColorText,
		},
		incomeLabel: &ui.Label{
			Text:     "$0.00",
			X:        centerX,
			Y:        ScreenHeight - 40,
			Color:    ColorText,
			Centered: true,
		},
		character: sprites.NewSmileyFace(),
	}

	cs.playButton.OnClick = onPlay

	return cs
}

func (cs *CoverScreen) Update() {
	cs.playButton.Update()
}

func (cs *CoverScreen) Draw(screen *ebiten.Image, face text.Face) {
	cs.levelLabel.Draw(screen, face)
	cs.bestScoreLabel.Draw(screen, face)
	cs.pickCharLabel.Draw(screen, face)
	cs.incomeLabel.Draw(screen, face)
	cs.playButton.Draw(screen, face)

	charW := cs.character.Bounds().Dx()
	charH := cs.character.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(ScreenWidth)/2-float64(charW)/2,
		float64(ScreenHeight)/2-float64(charH)/2+20,
	)
	screen.DrawImage(cs.character, op)
}
