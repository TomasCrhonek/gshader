package main

import (
	"testing"
)

func BenchmarkComputeMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		computeImage(1024, 1024, shaderMandelbrot, rainbow)
	}
}
