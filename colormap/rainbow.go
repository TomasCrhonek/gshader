package colormap

import (
	"image/color"
	"math"
)

func grayToRad(x uint8) float64 {
	return (float64(x) / 255.0) * 2 * math.Pi
}

func offset(x float64) uint8 {
	return uint8(127*x + 127)
}

func Rainbow(c color.Gray) color.Color {
	return color.RGBA{
		R: offset(math.Sin(grayToRad(c.Y))),
		G: offset(math.Sin(grayToRad(c.Y) + math.Pi/2)),
		B: offset(math.Sin(grayToRad(c.Y) + math.Pi)),
		A: 255,
	}
}
