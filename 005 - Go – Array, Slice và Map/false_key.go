package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10

	value, ok := x["key"]
	fmt.Println(value, ok)

	value2, ok2 := x["key2"]
	fmt.Println(value2, ok2)

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements)
}
