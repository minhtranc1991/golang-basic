package main

import "fmt"

func main() {
	var i int = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
