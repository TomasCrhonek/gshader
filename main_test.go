package main

import (
	"testing"

	"heronovo.cz/gshader/colormap"
	"heronovo.cz/gshader/shader"
)

func BenchmarkComputeMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeImage(1024, 1024, shader.Mandelbrot, colormap.Rainbow)
	}
}
