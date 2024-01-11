package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("test.txt")
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

	fmt.Printf("data sample size: %v\n", len(xs))

	mean := calculateMean(xs)

	sort.Float64s(xs)
	median := calculateMedian(xs)

	variance, stddev := calculateVarianceAndStdDev(xs, mean)

	fmt.Printf("mean=     %v\n", mean)
	fmt.Printf("median=   %v\n", median)
	fmt.Printf("variance= %v\n", variance)
	fmt.Printf("std-dev=  %v\n", stddev)
}

func calculateMean(xs []float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func calculateMedian(xs []float64) float64 {
	n := len(xs)
	if n%2 == 0 {
		return (xs[n/2-1] + xs[n/2]) / 2
	}
	return xs[n/2]
}

func calculateVarianceAndStdDev(xs []float64, mean float64) (float64, float64) {
	var sumSquaredDiff float64
	for _, x := range xs {
		diff := x - mean
		sumSquaredDiff += diff * diff
	}
	variance := sumSquaredDiff / float64(len(xs))
	stddev := math.Sqrt(variance)
	return variance, stddev
}
