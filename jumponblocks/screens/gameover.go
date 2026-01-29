package screens

import (
	"fmt"
	"image/color"

	"jumponblocks/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameOverScreen struct {
	earnings        float64
	titleLabel      *ui.Label
	earningsLabel   *ui.Label
	playAgainButton *ui.Button
	menuButton      *ui.Button
}

func NewGameOverScreen(onPlayAgain, onMenu func()) *GameOverScreen {
	centerX := float64(ScreenWidth) / 2

	gos := &GameOverScreen{
		titleLabel: &ui.Label{
			Text:     "GAME OVER",
			X:        centerX,
			Y:        120,
			Color:    color.White,
			Centered: true,
		},
		earningsLabel: &ui.Label{
			Text:     "Earned: $0.00",
			X:        centerX,
			Y:        200,
			Color:    color.RGBA{255, 215, 0, 255}, // Gold
			Centered: true,
		},
		playAgainButton: ui.NewButton(
			"PLAY AGAIN",
			centerX-70, 280,
			140, 50,
			ColorButton,
			ColorButtonHover,
			ColorText,
		),
		menuButton: ui.NewButton(
			"MENU",
			centerX-70, 350,
			140, 50,
			color.RGBA{100, 100, 100, 255},
			color.RGBA{130, 130, 130, 255},
			ColorText,
		),
	}

	gos.playAgainButton.OnClick = onPlayAgain
	gos.menuButton.OnClick = onMenu

	return gos
}

func (gos *GameOverScreen) SetScore(score int) {
	gos.earnings = float64(score) * 0.10
	gos.earningsLabel.Text = fmt.Sprintf("Earned: $%.2f", gos.earnings)
}

func (gos *GameOverScreen) Update() {
	gos.playAgainButton.Update()
	gos.menuButton.Update()
}

func (gos *GameOverScreen) Draw(screen *ebiten.Image, face text.Face) {
	gos.titleLabel.Draw(screen, face)
	gos.earningsLabel.Draw(screen, face)
	gos.playAgainButton.Draw(screen, face)
	gos.menuButton.Draw(screen, face)
}
