package mathematician

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) ([]float64, error) {
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
