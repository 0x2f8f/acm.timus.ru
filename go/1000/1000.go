package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(a + b)
}