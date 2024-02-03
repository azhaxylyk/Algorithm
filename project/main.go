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

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var xs []float64
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		var v float64
		txt := scan.Text()
		_, err = fmt.Sscanf(txt, "%f", &v)
		if err != nil {
			log.Fatalf(
				"could not convert to float64 %q: %v",
				txt, err,
			)
		}
		xs = append(xs, v)
	}

	if err = scan.Err(); err != nil {
		log.Fatalf("error scanning file: %v", err)
	}

	if len(xs) == 0 {
		fmt.Println("Error: Empty slice. Please provide a file with numeric values.")
		return
	}

	mean := stat.Mean(xs, nil)
	variance := stat.Variance(xs, nil)
	stddev := math.Sqrt(variance)

	sort.Float64s(xs)
	median := stat.Quantile(0.5, stat.Empirical, xs, nil)

	fmt.Printf("Average: %.0f\n", mean)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stddev)
}
