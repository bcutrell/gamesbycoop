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
	ebiten.SetTPS(60)                       // 60 TPS for smooth physics
	ebiten.SetRunnableOnUnfocused(false)    // Pause when window not focused
	ebiten.SetScreenClearedEveryFrame(true) // Explicit screen clearing

	g := game.NewGame()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
