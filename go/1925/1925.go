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

	var n, k, b, g int
	csum:=0
	sum:=0
	fmt.Scan(&n)
	fmt.Scan(&k)
	for i:=0; i<n; i++ {
		fmt.Scan(&b)
		fmt.Scan(&g)
		sum+=b
		csum+=g
	}
	sum=sum-n*2+k
	csum+=2
	if sum>=csum {
		fmt.Println(sum-csum)
	} else {
		fmt.Println("Big Bang!")
	}
}