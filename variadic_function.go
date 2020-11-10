package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	add1 := func(args ...int) int {
		total := 0
		for _, v := range args {
			total += v
		}
		return total
	}
	fmt.Println(add1(1, 2, 3, 4))
}
