package main

import "fmt"

type fraction struct {
	arg1, arg2 int
}

func (e *fraction) Error() string {
	return fmt.Sprintf("%d can't divide by %d", e.arg1, e.arg2)
}

func divide(arg1, arg2 int) (int, error) {
	if arg2 == 0 {
		err := fraction{arg1, arg2}
		return -1, &err
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
