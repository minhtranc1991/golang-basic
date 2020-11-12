package main

import "fmt"
import "sort"

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

//func (this ByName) Less(i, j int) bool {
//	return this[i].Name < this[j].Name
//}

func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this ByName) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}

func main() {
	kids := []Person{
		{"Jill", 19},
		{"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
