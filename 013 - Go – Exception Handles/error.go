package main

import (
	"errors"
	"fmt"
)

func divide(arg1, arg2 int) (int, error) {
	if arg2 == 0 {
		return -1, errors.New("Can't divide by 0")
	}
	return arg1 / arg2, nil
}

func main() {
	arg1 := 10
	arg2 := 0
	result, err := divide(arg1, arg2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
