package main

import (
	"fmt"
	"math"
	"os"
    Mt "mathematician/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your-program.go data.txt")
		return
	}
	filename := os.Args[1]
	nbrs, err := Mt.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if len(nbrs) == 0 {
		fmt.Println("No data to process.")
		return
	}
	avg := Mt.Average(nbrs)
	med := Mt.Median(nbrs)
	hh := Mt.Variance(nbrs, avg)
	stdDev :=  Mt.StandardDeviation(hh)
	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	fmt.Printf("Median: %d\n", int(math.Round(med)))
	fmt.Printf("Variance: %d\n", int(math.Round(hh)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
