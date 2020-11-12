package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello server"
	fmt.Println("Sending : ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	_ = c.Close()
}

func main() {
	client()
}
