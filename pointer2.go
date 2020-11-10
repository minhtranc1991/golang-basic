package main

import "fmt"

func zero(xPointer *int) {
	*xPointer = 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
