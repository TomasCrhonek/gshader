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

func mono(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	return uint8((0.2125 * float64(r)) + (0.7154 * float64(g)) + (0.0721 * float64(b)))
}
