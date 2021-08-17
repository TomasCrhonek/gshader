package main

import "image/color"

func shaderGradient(x, y, w, h int) color.Gray {
	return color.Gray{uint8(255 * x / w)}
}
