package main

import (
	"testing"

	"github.com/TomasCrhonek/gshader/colormap"
	"github.com/TomasCrhonek/gshader/shader"
)

func BenchmarkComputeMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeImage(1024, 1024, shader.Mandelbrot, colormap.Rainbow)
	}
}

func BenchmarkComputeJulia1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeImage(1024, 1024, shader.Julia1, colormap.Rainbow)
	}
}

func BenchmarkComputeJulia2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeImage(1024, 1024, shader.Julia2, colormap.Rainbow)
	}
}

func BenchmarkComputeRandomNoise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeImage(1024, 1024, shader.RandomNoise, colormap.Rainbow)
	}
}

func BenchmarkComputeGradient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeImage(1024, 1024, shader.Gradient, colormap.Rainbow)
	}
}
