package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"runtime/trace"
)

const (
	DEFAULT_IMAGEFILENAME = "image.png"
	DEFAULT_WIDTH         = 1024
	DEFAULT_HEIGHT        = 1024
	DEFAULT_TRACEFILENAME = "main.trace"
)

const help string = `
Usage: gshader [-h] [-t] [-x pixels] [-y pixels] [-n filename]

  -x width  Image width in pixels.
  -y height Image height in pixels.
  -n name   Image file name.

  -t        Generate trace information about performance. Use with 'go tool trace main.trace'.

  -h        This help page.
`

func main() {
	traceQ := flag.Bool("t", false, "")
	helpQ := flag.Bool("h", false, "")
	imageFile := flag.String("n", DEFAULT_IMAGEFILENAME, "")
	width := flag.Int("x", DEFAULT_WIDTH, "")
	height := flag.Int("y", DEFAULT_HEIGHT, "")
	flag.Usage = func() { fmt.Println(help) }
	flag.Parse()

	if *helpQ {
		flag.Usage()
		return
	}

	if *traceQ {
		traceFile, err := os.Create(DEFAULT_TRACEFILENAME)
		if err != nil {
			log.Fatal(err)
		}
		defer traceFile.Close()
		trace.Start(traceFile)
		defer trace.Stop()
	}

	f, err := os.Create(*imageFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	encoder := png.Encoder{CompressionLevel: png.NoCompression}
	if err := encoder.Encode(f, computeImage(*width, *height, shaderMandelbrot)); err != nil {
		log.Fatal(err)
	}
}
