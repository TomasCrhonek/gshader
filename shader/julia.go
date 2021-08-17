package shader

import (
	"image/color"
	"math/cmplx"
)

func Julia1(x int, y int, width int, height int) color.Color {
	z := complex(normalize(x, width, -2, 2), normalize(y, height, -2, 2))

	const ITER = 32
	c := -1 + 0i

	iter := 0

	for ; (cmplx.Abs(z) <= 2) && (iter < ITER); iter++ {
		z = z*z + c
	}

	return color.Gray{uint8((255 * (iter)) / ITER)}
}

func Julia2(x int, y int, width int, height int) color.Color {
	z := complex(normalize(x, width, -2, 2), normalize(y, height, -2, 2))

	const ITER = 32
	c := 0.3 + 0.6i

	iter := 0

	for ; (cmplx.Abs(z) <= 2) && (iter < ITER); iter++ {
		z = z*z + c
	}

	return color.Gray{uint8((255 * (iter)) / ITER)}
}
