package screens

import (
	"fmt"
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
	pickCharButton *ui.Button
	incomeLabel    *ui.Label
	character      *ebiten.Image
	selectedChar   sprites.CharacterType
	wallet         *float64
}

func NewCoverScreen(onPlay func(), onCharSelect func()) *CoverScreen {
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
		pickCharButton: ui.NewButton(
			"Player",
			ScreenWidth-100, ScreenHeight-95,
			80, 30,
			ColorButton,
			ColorButtonHover,
			ColorText,
		),
		incomeLabel: &ui.Label{
			Text:     "$0.00",
			X:        centerX,
			Y:        ScreenHeight - 40,
			Color:    ColorText,
			Centered: true,
		},
		character:    sprites.NewSmileyFace(),
		selectedChar: sprites.CharacterSmiley,
	}

	cs.playButton.OnClick = onPlay
	cs.pickCharButton.OnClick = onCharSelect

	return cs
}

func (cs *CoverScreen) SetWallet(wallet *float64) {
	cs.wallet = wallet
}

func (cs *CoverScreen) SetBestScore(best int) {
	earnings := float64(best) * 0.10
	cs.bestScoreLabel.Text = fmt.Sprintf("Best: $%.2f", earnings)
}

func (cs *CoverScreen) SetSelectedCharacter(charType sprites.CharacterType) {
	cs.selectedChar = charType
	cs.character = sprites.NewCharacter(charType)
}

func (cs *CoverScreen) GetSelectedCharacter() sprites.CharacterType {
	return cs.selectedChar
}

func (cs *CoverScreen) Update() {
	cs.playButton.Update()
	cs.pickCharButton.Update()

	// Update income label from wallet
	if cs.wallet != nil {
		cs.incomeLabel.Text = fmt.Sprintf("$%.2f", *cs.wallet)
	}
}

func (cs *CoverScreen) Draw(screen *ebiten.Image, face text.Face) {
	cs.levelLabel.Draw(screen, face)
	cs.bestScoreLabel.Draw(screen, face)
	cs.pickCharButton.Draw(screen, face)
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
