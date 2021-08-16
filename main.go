package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"sync"
)

const (
	imgFile = "image.png"
	width   = 10240
	height  = 10240
)

func main() {
	f, err := os.Create(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, computeImage(width, height)); err != nil {
		log.Fatal(err)
	}
}

func computeImage(width, height int) image.Image {
	img := image.NewGray(image.Rect(0, 0, width, height))

	var wg sync.WaitGroup

	wg.Add(width)

	for col := 0; col < width; col++ {
		go func(r int) {
			for row := 0; row < height; row++ {
				img.Set(r, row, shaderMandelbrot(r, row, width, height))
			}
			wg.Done()
		}(col)
	}

	wg.Wait()
	return img
}
