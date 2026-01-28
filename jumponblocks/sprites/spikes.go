package sprites

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var ColorSpike = color.RGBA{200, 50, 50, 255} // Red
var ColorSpikeDark = color.RGBA{150, 30, 30, 255} // Darker red for pattern

// DrawSpikes draws spikes around the border of the screen
func DrawSpikes(screen *ebiten.Image, screenWidth, screenHeight, thickness int) {
	t := float32(thickness)
	w := float32(screenWidth)
	h := float32(screenHeight)

	// Top edge
	vector.FillRect(screen, 0, 0, w, t, ColorSpike, false)
	// Bottom edge
	vector.FillRect(screen, 0, h-t, w, t, ColorSpike, false)
	// Left edge
	vector.FillRect(screen, 0, 0, t, h, ColorSpike, false)
	// Right edge
	vector.FillRect(screen, w-t, 0, t, h, ColorSpike, false)

	// Draw stripe pattern to indicate danger
	stripeWidth := float32(8)

	// Top stripes
	for x := float32(0); x < w; x += stripeWidth * 2 {
		vector.FillRect(screen, x, 0, stripeWidth, t, ColorSpikeDark, false)
	}

	// Bottom stripes
	for x := float32(0); x < w; x += stripeWidth * 2 {
		vector.FillRect(screen, x, h-t, stripeWidth, t, ColorSpikeDark, false)
	}

	// Left stripes
	for y := float32(0); y < h; y += stripeWidth * 2 {
		vector.FillRect(screen, 0, y, t, stripeWidth, ColorSpikeDark, false)
	}

	// Right stripes
	for y := float32(0); y < h; y += stripeWidth * 2 {
		vector.FillRect(screen, w-t, y, t, stripeWidth, ColorSpikeDark, false)
	}
}
