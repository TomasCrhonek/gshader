package main

import (
	"image"
	"image/color"
	"sync"
)

type shader func(int, int, int, int) color.Gray
type colorMap func(color.Gray) color.Color

func computeImage(width, height int, shaderFunc shader, colorFunc colorMap) image.Image {
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
