package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"os"
	"runtime/trace"

	"heronovo.cz/gshader/colormap"
	"heronovo.cz/gshader/shader"
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

  -l 		List avaiable shaders and colormaps.
  -s shader Select shader. List shaders with -l.
  -c color  Select colormap. List colormaps with -l.

  -t        Generate trace information about performance. Use with 'go tool trace main.trace'.

  -h        This help page.
`

func main() {
	traceQ := flag.Bool("t", false, "")
	helpQ := flag.Bool("h", false, "")
	imageFile := flag.String("n", DEFAULT_IMAGEFILENAME, "")
	width := flag.Int("x", DEFAULT_WIDTH, "")
	height := flag.Int("y", DEFAULT_HEIGHT, "")
	list := flag.Bool("l", false, "")
	shd := flag.String("s", "", "")
	clr := flag.String("c", "", "")
	flag.Usage = func() { fmt.Println(help) }
	flag.Parse()

	if *helpQ {
		flag.Usage()
		return
	}

	registerShader(shader.Mandelbrot, "mandel")
	registerShader(shader.Gradient, "gradient")
	registerShader(shader.Julia1, "julia1")
	registerShader(shader.Julia2, "julia2")
	registerShader(shader.RandomNoise, "randomnoise")

	registerColormap(colormap.Red, "red")
	registerColormap(colormap.Rainbow, "rainbow")

	if *list {
		listShadersColormaps()
		return
	}

	var shaderFunc shader.Shader
	if len(*shd) > 0 {
		shaderFunc = getShader(*shd)
		if shaderFunc == nil {
			fmt.Fprintf(os.Stderr, "Shader '%s' not found.\n", *shd)
			os.Exit(1)
		}
	} else {
		shaderFunc = shaderList[0].shader
	}

	var colorFunc colormap.ColorMap
	if len(*clr) > 0 {
		colorFunc = getColormap(*clr)
		if shaderFunc == nil {
			fmt.Fprintf(os.Stderr, "ColorMap '%s' not found.\n", *clr)
			os.Exit(1)
		}
	} else {
		colorFunc = colorList[0].color
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
	if err := encoder.Encode(f, computeImage(*width, *height, shaderFunc, colorFunc)); err != nil {
		log.Fatal(err)
	}
}
