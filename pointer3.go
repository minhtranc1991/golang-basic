package main

import "fmt"

func main() {
	var x *int
	var y int
	y = 0
	x = &y

	fmt.Println(x)
	fmt.Println(&y)
	fmt.Println(*x)
	fmt.Println(y)
	fmt.Println(&x)

	*x = 1

	fmt.Println(*x)
	fmt.Println(y)
}
