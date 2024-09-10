package mathematician



func Average(nbrs []float64) float64 {
    total := 0.0
    for _, num := range nbrs {
        total += num
    }
    return total / float64(len(nbrs))
}

