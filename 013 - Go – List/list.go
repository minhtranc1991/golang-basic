package main

import "fmt"
import "container/list"

func main() {
	var x list.List

	x.PushBack(1)
	x.PushBack(2.5)
	x.PushBack("Hello")

	for e := x.Front() ; e != nil ; e = e.Next() {
		fmt.Println(e.Value)
	}
}