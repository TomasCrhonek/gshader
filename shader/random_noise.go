package shader

import (
	"image/color"
	"math/rand"
)

func RandomNoise(x int, y int, width int, height int) color.Color {
	return color.Gray{uint8(rand.Intn(256))}
}
