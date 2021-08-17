package shader

import "testing"

func BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mandelbrot(1920/2, 1080/2, 1920, 1080)
	}
}
