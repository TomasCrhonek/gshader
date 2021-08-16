package main

import (
	"image/png"
	"log"
	"os"
	"runtime/trace"

	flag "github.com/spf13/pflag"
)

const (
	DEFAULT_IMAGEFILENAME = "image.png"
	DEFAULT_WIDTH         = 1024
	DEFAULT_HEIGHT        = 1024
)

func main() {
	traceQ := flag.BoolP("trace", "t", false, "Trace the program for performance analysis.")
	helpQ := flag.BoolP("help", "h", false, "This help page.")
	imageFile := flag.StringP("name", "n", DEFAULT_IMAGEFILENAME, "Image `file` name.")
	width := flag.IntP("width", "x", DEFAULT_WIDTH, "Image `width`.")
	height := flag.IntP("height", "y", DEFAULT_HEIGHT, "Image `height`.")
	flag.CommandLine.SortFlags = false
	flag.Parse()

	if *helpQ {
		flag.Usage()
		return
	}

	if *traceQ {
		traceFile, err := os.Create("main.trace")
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
