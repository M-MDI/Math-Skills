package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run program.go <filename>")
		return
	}

	filename := os.Args[1]

	data, err := readData(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	avg := calculateAverage(data)
	median := calculateMedian(data)
	variance := calculateVariance(data)
	stdDev := calculateStandardDeviation(variance, len(data))

	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Median: %.2f\n", median)
	fmt.Printf("Variance: %.2f\n", variance)
	fmt.Printf("Standard Deviation: %.2f\n", stdDev)
}

func readData(filename string) ([]float64, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var numbers []float64
	for _, line := range strings.Split(string(data), "\n") {
		num, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			return nil, fmt.Errorf("invalid number in file: %w", err)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

func calculateMedian(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 1 {
		return data[n/2]
	} else {
		return (data[n/2-1] + data[n/2]) / 2
	}
}

func calculateVariance(data []float64) float64 {
	mean := calculateAverage(data)
	variance := 0.0
	for _, v := range data {
		variance += math.Pow(v-mean, 2)
	}
	return variance / float64(len(data)-1)
}

func calculateStandardDeviation(variance float64, n int) float64 {
	return math.Sqrt(variance)
}
