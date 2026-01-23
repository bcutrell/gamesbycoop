package game

type GameState int

const (
	StateCover GameState = iota
	StateFlash
)

const FlashDuration = 90 // frames (1.5 seconds at 60 FPS)
