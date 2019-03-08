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

	var n, k, v int
	fmt.Scan(&n)
	fmt.Scan(&k)

	d := 0
	g := 0

	for i:=0; i<n; i++ {
		fmt.Scan(&v)
		if v>k {
			g+=v-k
		} else {
			d+=k-v
		}
	}
	fmt.Println(g, d)
}