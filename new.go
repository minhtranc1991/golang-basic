package main

import "fmt"

func one(xPointer *int) {
	*xPointer = 1
}

func main() {
	xPointer := new(int)
	one(xPointer)
	fmt.Println(*xPointer)
}
