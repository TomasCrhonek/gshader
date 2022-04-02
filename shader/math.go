package shader

func normalize(x int, total int, min float64, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
