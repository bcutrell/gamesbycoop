package main

import (
	"log"
	"jumponblocks/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle("Jump on Blocks")
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(30)                       // Lower tick rate to reduce GPU load
	ebiten.SetRunnableOnUnfocused(false)    // Pause when window not focused
	ebiten.SetScreenClearedEveryFrame(true) // Explicit screen clearing

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
