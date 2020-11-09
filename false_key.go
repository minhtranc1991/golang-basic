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
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be" : "Beryllium",
		"B" : "Boron",
		"C" : "Carbon",
		"N" : "Nitrogen",
		"O" : "Oxygen",
		"F" : "Fluorine",
		"Ne" : "Neon",
	}

	elements2 := make(map[string]map[string]string) {
        "H" : map[string]string {
        "name" : "Hydrogen",
    },
        "He" : map[string]string {
        "name" : "Helium",
    },
        "Li" : map[string]string {
        "name" : "Lithium",
    },
        "Be" : map[string]string {
        "name" : "Beryllium",
    },
        "B" : map[string]string {
         "name" : "Boron",
    },
        "C" : map[string]string {
        "name" : "Carbon",
    },
        "N" : map[string]string {
        "name" : "Nitrogen",
    },
        "O" : map[string]string {
        "name" : "Oxygen",
    },
        "F" : map[string]string {
        "name" : "Fluorine",
    },
        "Ne" : map[string]string {
        "name" : "Neon",
    }, 
		}
	}