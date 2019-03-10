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

	var n int
	fmt.Scan(&n)
	switch n {
	case 1:
		fmt.Println(1,2,3)
	case 2:
		fmt.Println(3,4,5)
	default:
		fmt.Println(-1)
	}
}