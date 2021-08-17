package colormap

import (
	"image/color"
	"testing"
)

func BenchmarkRainbow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Rainbow(color.Gray{127})
	}
}
