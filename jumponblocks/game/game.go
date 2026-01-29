package game

import (
	"fmt"
	"image/color"

	"jumponblocks/screens"
	"jumponblocks/sprites"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	state            GameState
	flashTimer       int
	coverScreen      *screens.CoverScreen
	charSelectScreen *screens.CharSelectScreen
	gameplayScreen   *screens.GameplayScreen
	gameOverScreen   *screens.GameOverScreen
	font             text.Face
	clickLock        bool
	bestScore        int
	wallet           float64
}

func NewGame() *Game {
	g := &Game{state: StateCover, wallet: 0.00} // Start with $0
	g.font = text.NewGoXFace(basicfont.Face7x13)

	// Set up gameplay screen first so we can reference it
	g.gameplayScreen = screens.NewGameplayScreen(func(score int) {
		if score > g.bestScore {
			g.bestScore = score
			g.coverScreen.SetBestScore(score)
		}
		g.wallet += float64(score) * 0.10 // $0.10 per second survived
		g.state = StateGameOver
		g.gameOverScreen.SetScore(score)
	})

	// Set up cover screen
	g.coverScreen = screens.NewCoverScreen(
		func() {
			// Play button
			if !g.clickLock {
				g.state = StateFlash
				g.flashTimer = FlashDuration
				g.clickLock = true
			}
		},
		func() {
			// Character select button
			if !g.clickLock {
				g.charSelectScreen.SetSelectedChar(g.coverScreen.GetSelectedCharacter())
				g.state = StateCharSelect
				g.clickLock = true
			}
		},
	)

	// Connect wallet to cover screen
	g.coverScreen.SetWallet(&g.wallet)

	// Set up character selection screen
	g.charSelectScreen = screens.NewCharSelectScreen(
		&g.wallet,
		func() {
			// Back button
			g.state = StateCover
		},
		func(charType sprites.CharacterType) {
			// Character selected
			g.gameplayScreen.SetCharacter(charType)
			g.coverScreen.SetSelectedCharacter(charType)
		},
	)

	// Set up game over screen
	g.gameOverScreen = screens.NewGameOverScreen(
		func() {
			// Play again
			if !g.clickLock {
				g.state = StateFlash
				g.flashTimer = FlashDuration
				g.clickLock = true
			}
		},
		func() {
			// Back to menu
			if !g.clickLock {
				g.state = StateCover
				g.clickLock = true
			}
		},
	)

	return g
}

func (g *Game) Update() error {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.clickLock = false
	}

	switch g.state {
	case StateCover:
		g.coverScreen.Update()
	case StateCharSelect:
		g.charSelectScreen.Update()
	case StateFlash:
		g.flashTimer--
		if g.flashTimer <= 0 {
			g.gameplayScreen.Reset()
			g.state = StatePlaying
		}
	case StatePlaying:
		g.gameplayScreen.Update()
	case StateGameOver:
		g.gameOverScreen.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateCover:
		screen.Fill(ColorBackground)
		g.coverScreen.Draw(screen, g.font)
	case StateCharSelect:
		screen.Fill(ColorBackground)
		g.charSelectScreen.Draw(screen, g.font)
	case StateFlash:
		screen.Fill(ColorFlashBg)
		// Show countdown: 3, 2, 1, Go!
		var msg string
		remaining := g.flashTimer
		if remaining > 135 { // 180-135 = first 45 frames
			msg = "3"
		} else if remaining > 90 { // 135-90 = next 45 frames
			msg = "2"
		} else if remaining > 45 { // 90-45 = next 45 frames
			msg = "1"
		} else { // last 45 frames
			msg = "Go!"
		}
		op := &text.DrawOptions{}
		w, h := text.Measure(msg, g.font, 0)
		op.GeoM.Translate(float64(ScreenWidth)/2-w/2, float64(ScreenHeight)/2-h/2)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, msg, g.font, op)
	case StatePlaying:
		screen.Fill(ColorBackground)
		g.gameplayScreen.Draw(screen)
		// Draw earnings at top center
		earnings := float64(g.gameplayScreen.GetScore()) * 0.10
		earningsMsg := fmt.Sprintf("$%.2f", earnings)
		op := &text.DrawOptions{}
		sw, _ := text.Measure(earningsMsg, g.font, 0)
		op.GeoM.Translate(float64(ScreenWidth)/2-sw/2, 30)
		op.ColorScale.ScaleWithColor(color.RGBA{255, 215, 0, 255}) // Gold color
		text.Draw(screen, earningsMsg, g.font, op)
	case StateGameOver:
		screen.Fill(ColorFlashBg)
		g.gameOverScreen.Draw(screen, g.font)
		// Show best earnings
		bestEarnings := float64(g.bestScore) * 0.10
		bestMsg := fmt.Sprintf("Best: $%.2f", bestEarnings)
		op := &text.DrawOptions{}
		bw, _ := text.Measure(bestMsg, g.font, 0)
		op.GeoM.Translate(float64(ScreenWidth)/2-bw/2, 240)
		op.ColorScale.ScaleWithColor(color.RGBA{200, 200, 200, 255})
		text.Draw(screen, bestMsg, g.font, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
