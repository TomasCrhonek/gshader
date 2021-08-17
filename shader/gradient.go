package shader

import "image/color"

func Gradient(x, y, w, h int) color.Color {
	return color.Gray{uint8(255 * x / w)}
}
