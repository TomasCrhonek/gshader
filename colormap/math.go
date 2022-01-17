package colormap

import "math"

func grayToRad(x uint8) float64 {
	return (float64(x) / 255.0) * 2 * math.Pi
}

func offset(x float64) uint8 {
	return uint8(127*x + 127)
}
