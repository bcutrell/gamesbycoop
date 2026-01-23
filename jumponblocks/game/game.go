package game

import (
	"image/color"

	"jumponblocks/screens"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	state       GameState
	flashTimer  int
	coverScreen *screens.CoverScreen
	font        text.Face
	clickLock   bool
}

func NewGame() *Game {
	g := &Game{state: StateCover}
	g.font = text.NewGoXFace(basicfont.Face7x13)
	g.coverScreen = screens.NewCoverScreen(func() {
		if !g.clickLock {
			g.state = StateFlash
			g.flashTimer = FlashDuration
			g.clickLock = true
		}
	})
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
			g.state = StateCover
		}
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
		msg := "GAME STARTS"
		op := &text.DrawOptions{}
		w, h := text.Measure(msg, g.font, 0)
		op.GeoM.Translate(float64(ScreenWidth)/2-w/2, float64(ScreenHeight)/2-h/2)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(screen, msg, g.font, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
