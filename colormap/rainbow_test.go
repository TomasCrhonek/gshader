package colormap

import (
	"image/color"
	"testing"
)

func BenchmarkRainbow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 256; i++ {
			Rainbow(color.Gray{uint8(i)})
		}
	}
}
