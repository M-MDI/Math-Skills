package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func readNumbersFromFile(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, scanner.Err()
}

func calculateAverage(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func calculateVariance(numbers []int, mean float64) float64 {
	var sumSquares float64
	for _, num := range numbers {
		diff := float64(num) - mean
		sumSquares += diff * diff
	}
	return sumSquares / float64(len(numbers))
}

func calculateStdDev(variance float64) float64 {
	return math.Sqrt(variance)
}

func writeResultsToFile(filePath string, avg, median, variance, stdDev float64) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Average: %.2f\nMedian: %.2f\nVariance: %.2f\nStandard Deviation: %.2f\n", avg, median, variance, stdDev)
	return err
}

func calculateMedian(numbers []int) float64 {
	n := len(numbers)
	mid := n / 2
	if n%2 == 0 {
		return float64(numbers[mid-1]+numbers[mid]) / 2
	}
	return float64(numbers[mid])
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the file path as an argument")
	}
	filePath := os.Args[1]

	numbers, err := readNumbersFromFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Sort numbers to calculate the median
	sort.Ints(numbers)

	average := calculateAverage(numbers)
	median := calculateMedian(numbers)
	variance := calculateVariance(numbers, average)
	stdDev := calculateStdDev(variance)

	err = writeResultsToFile("statistics.txt", average, median, variance, stdDev)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Statistics written to statistics.txt")
}
