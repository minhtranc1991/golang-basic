package main

import "fmt"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)
	fmt.Println(a, ",", b, ",", c)
}
