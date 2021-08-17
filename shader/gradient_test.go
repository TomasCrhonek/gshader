package shader

import "testing"

func BenchmarkGradient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gradient(1920/2, 1080/2, 1920, 1080)
	}
}
