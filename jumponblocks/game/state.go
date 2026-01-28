package game

type GameState int

const (
	StateCover GameState = iota
	StateFlash
	StatePlaying
	StateGameOver
)

const FlashDuration = 180 // frames (3 seconds at 60 FPS for countdown)
