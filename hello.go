package main

import (
	"fmt"
	"os"
)

func main() {
	length := len(os.Args)
	fmt.Printf("length=%d\n", length)
	fmt.Printf("hello, world!!!!!!\n")
	if length >= 2 {
		fmt.Println("command=" + os.Args[1])
	}
	fmt.Println(os.Args)
}
