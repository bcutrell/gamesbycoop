package game

import (
	"fmt"
	"image/color"

	"jumponblocks/screens"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	state          GameState
	flashTimer     int
	coverScreen    *screens.CoverScreen
	gameplayScreen *screens.GameplayScreen
	gameOverScreen *screens.GameOverScreen
	font           text.Face
	clickLock      bool
	bestScore      int
	lastScore      int
}

func NewGame() *Game {
	g := &Game{state: StateCover}
	g.font = text.NewGoXFace(basicfont.Face7x13)

	// Set up cover screen
	g.coverScreen = screens.NewCoverScreen(func() {
		if !g.clickLock {
			g.state = StateFlash
			g.flashTimer = FlashDuration
			g.clickLock = true
		}
	})

	// Set up gameplay screen
	g.gameplayScreen = screens.NewGameplayScreen(func(score int) {
		g.lastScore = score
		if score > g.bestScore {
			g.bestScore = score
		}
		g.state = StateGameOver
		g.gameOverScreen.SetScore(score)
	})

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
		// Draw score at top center
		scoreMsg := fmt.Sprintf("Score: %d", g.gameplayScreen.GetScore())
		op := &text.DrawOptions{}
		sw, _ := text.Measure(scoreMsg, g.font, 0)
		op.GeoM.Translate(float64(ScreenWidth)/2-sw/2, 30)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, scoreMsg, g.font, op)
	case StateGameOver:
		screen.Fill(ColorFlashBg)
		g.gameOverScreen.Draw(screen, g.font)
		// Show best score
		bestMsg := fmt.Sprintf("Best: %d", g.bestScore)
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
