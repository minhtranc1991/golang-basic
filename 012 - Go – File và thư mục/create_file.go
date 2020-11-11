package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	_, _ = file.WriteString("test")
}
