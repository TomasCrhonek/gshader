package colormap

import "image/color"

func Red(c color.Color) color.Color {
	return color.RGBA{R: mono(c), A: 255}
}
