package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"gonum.org/v1/gonum/stat"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go data.txt")
		os.Exit(1)
	}

	fileName := os.Args[1]
	xs, err := readDataFromFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if len(xs) == 0 {
		fmt.Println("Error: No valid numeric values provided in the file.")
		return
	}

	mean := calculateMean(xs)
	variance := calculateVariance(xs, mean)
	stddev := math.Sqrt(variance)
	median := calculateMedian(xs)

	fmt.Printf("Average: %.0f\n", mean)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stddev)
}

func readDataFromFile(fileName string) ([]float64, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xs []float64
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		var v float64
		txt := scan.Text()
		if _, err = fmt.Sscanf(txt, "%f", &v); err != nil {
			log.Printf("Warning: could not convert %q to float64: %v\n", txt, err)
			continue
		}
		xs = append(xs, v)
	}

	if err = scan.Err(); err != nil {
		return nil, err
	}

	return xs, nil
}

func calculateMean(xs []float64) float64 {
	return stat.Mean(xs, nil)
}

func calculateVariance(xs []float64, mean float64) float64 {
	return stat.Variance(xs, nil)
}

func calculateMedian(xs []float64) float64 {
	sort.Float64s(xs)
	return stat.Quantile(0.5, stat.Empirical, xs, nil)
}
