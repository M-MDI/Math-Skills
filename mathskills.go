 
package main
import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
)
func readFile(filename string) ([]float64, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    var nbrs []float64
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if num, err := strconv.ParseFloat(line, 64); err == nil {
            nbrs = append(nbrs, num)
        }
    }
    return nbrs, scanner.Err()
}
func average(nbrs []float64) float64 {
    total := 0.0
    for _, num := range nbrs {
        total += num
    }
    return total / float64(len(nbrs))
}
func median(nbrs []float64) float64 {
    sort.Float64s(nbrs)
    n := len(nbrs)
    if n%2 == 1 {
        return nbrs[n/2]
    }
    return (nbrs[n/2-1] + nbrs[n/2]) / 2
}
func variance(nbrs []float64, mean float64) float64 {
    total := 0.0
    for _, num := range nbrs {
        diff := num - mean
        total += diff * diff
    }
    return total / float64(len(nbrs))
}
func standardDeviation(variance float64) float64 {
    return math.Sqrt(variance)
}
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run your-program.go data.txt")
        return
    }
    filename := os.Args[1]
    nbrs, err := readFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }
    if len(nbrs) == 0 {
        fmt.Println("No data to process.")
        return
    }
    avg := average(nbrs)
    med := median(nbrs)
    hh := variance(nbrs, avg)
    stdDev := standardDeviation(hh)
    fmt.Printf("Average: %d\n", int(math.Round(avg)))
    fmt.Printf("Median: %d\n", int(math.Round(med)))
    fmt.Printf("Variance: %d\n", int(math.Round(hh)))
    fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}