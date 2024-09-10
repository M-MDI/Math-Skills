package mathematician

import "sort"

func Median(nbrs []float64) float64 {
	sort.Float64s(nbrs)
	n := len(nbrs)
	if n%2 == 1 {
		return nbrs[n/2]
	}
	return (nbrs[n/2-1] + nbrs[n/2]) / 2
}
