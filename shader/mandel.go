package shader

import (
	"image/color"
	"math/cmplx"
)

func Mandelbrot(x int, y int, width int, height int) color.Color {
	c := complex(normalize(x, width, -2, 2), normalize(y, height, -2, 2))

	const INTER = 1024
	z := 0 + 0i

	iter := 0

	for ; (cmplx.Abs(z) <= 2) && (iter < INTER); iter++ {
		z = z*z + c
	}

	return color.Gray{uint8((255 * iter) / INTER)}
}

func normalize(x int, total int, min float64, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
