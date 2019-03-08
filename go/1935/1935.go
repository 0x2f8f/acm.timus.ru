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

	var n, v int
	max:=0
	sum:=0
	fmt.Scan(&n)
	for i:=0; i<n; i++ {
		fmt.Scan(&v)
		sum+=v
		if v>max {
			max=v
		}
	}
	sum+=max
	fmt.Println(sum)
}