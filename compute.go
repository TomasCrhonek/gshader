package main

import (
	"image"
	"sync"

	"github.com/TomasCrhonek/gshader/colormap"
	"github.com/TomasCrhonek/gshader/shader"
)

func computeImage(width, height int, shaderFunc shader.Shader, colorFunc colormap.ColorMap) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup

	wg.Add(width)

	for col := 0; col < width; col++ {
		go func(r int) {
			for row := 0; row < height; row++ {
				img.Set(r, row, colorFunc(shaderFunc(r, row, width, height)))
			}
			wg.Done()
		}(col)
	}

	wg.Wait()
	return img
}
