package colormap

import "image/color"

func Red(c color.Gray) color.Color {
	return color.RGBA{R: c.Y, A: 255}
}
