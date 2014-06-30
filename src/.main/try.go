package main

import (
	"fmt"
)

func main() {
	var b []byte
	fmt.Printf("len=%d cap=%d\n", len(b), cap(b))
	for i := 0; i < 8; i++ {
		b = append(b, "ami"...)
		fmt.Printf("len=%d cap=%d\n", len(b), cap(b))
	}
	
	fmt.Println(string(b))
}
