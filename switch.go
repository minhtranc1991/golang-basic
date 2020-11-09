package main

import "fmt"

func main() {
	var i int = 0
	for i <= 5 {
		if i == 0 {
			fmt.Println("Khong")
		} else if i == 1 {
			fmt.Println("Mot")
		} else if i == 2 {
			fmt.Println("Hai")
		} else if i == 3 {
			fmt.Println("Ba")
		} else if i == 4 {
			fmt.Println("Bon")
		} else if i == 5 {
			fmt.Println("Nam")
		}
		i++
	}
	i = 0
	for i <= 6 {
		switch i {
		case 0:
			fmt.Println("Khong")
		case 1:
			fmt.Println("Mot")
		case 2:
			fmt.Println("Hai")
		case 3:
			fmt.Println("Ba")
		case 4:
			fmt.Println("Bon")
		case 5:
			fmt.Println("Nam")
		default:
			fmt.Println("Khong biet")
		}
		i++
	}
}
