package colormap

import (
	"image/color"
	"math"
)

func Rainbow(c color.Color) color.Color {
	return color.RGBA{
		R: offset(math.Sin(grayToRad(mono(c)))),
		G: offset(math.Sin(grayToRad(mono(c)) + math.Pi/2)),
		B: offset(math.Sin(grayToRad(mono(c)) + math.Pi)),
		A: 255,
	}
}
