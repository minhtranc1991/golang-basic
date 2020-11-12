package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str1 := []byte("There is no one who loves pain itself, who seeks after it and wants to have it, simply because it is pain")
	crypto := sha1.New()
	crypto.Write(str1)
	fmt.Println(crypto.Sum([]byte{}))
	fmt.Printf("%x", crypto.Sum([]byte{}))
}
