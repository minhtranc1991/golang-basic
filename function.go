package main

import "fmt"

func average(input []float64) float64 {
	total := 0.0
	for _, v := range input {
		total += v
	}
	return total / float64(len(input))
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
}
