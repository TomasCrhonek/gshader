package main

import (
	"image/png"
	"log"
	"os"
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

	if err := png.Encode(f, computeImage(width, height, shaderMandelbrot)); err != nil {
		log.Fatal(err)
	}
}
