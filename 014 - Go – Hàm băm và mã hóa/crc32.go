package main

import (
	"fmt"
	"hash/crc32"
)

func main() {

	str1 := []byte("Tran Cong Minh")
	hash1 := crc32.NewIEEE()
	hash1.Write(str1)
	fmt.Println(hash1.Sum32())

	str2 := []byte("There is no one who loves pain itself, who seeks after it and wants to have it, simply because it is pain...")
	hash2 := crc32.NewIEEE()
	hash2.Write(str2)
	fmt.Println(hash2.Sum32())

	if hash1.Sum32() != hash2.Sum32() {
		fmt.Println("Different")
	} else {
		fmt.Println("Same")
	}
}
