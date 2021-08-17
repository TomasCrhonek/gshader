package colormap

import "image/color"

func mono(c color.Color) uint8 {
	r, g, b, _ := c.RGBA()
	return uint8((0.2125 * float64(r)) + (0.7154 * float64(g)) + (0.0721 * float64(b)))
}

func Red(c color.Color) color.Color {
	return color.RGBA{R: mono(c), A: 255}
}
