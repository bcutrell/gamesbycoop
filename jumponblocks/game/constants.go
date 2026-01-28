package game

import "image/color"

const (
	ScreenWidth  = 400
	ScreenHeight = 600
)

// Physics constants
const (
	Gravity        = 0.5
	JumpVelocity   = -12.0
	MoveSpeed      = 5.0
	BlockWidth     = 80
	BlockHeight    = 20
	SpikeThickness = 15
)

var (
	ColorBackground = color.RGBA{135, 206, 235, 255}
	ColorFlashBg    = color.RGBA{0, 0, 0, 255}
)
