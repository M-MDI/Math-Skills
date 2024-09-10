package mathematician

func Variance(nbrs []float64, mean float64) float64 {
	total := 0.0
	for _, num := range nbrs {
		diff := num - mean
		total += diff * diff
	}
	return total / float64(len(nbrs))
}
