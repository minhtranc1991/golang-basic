package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
