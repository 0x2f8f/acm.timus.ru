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

	var n, m, y, x, l, i, j uint64
	f := 0
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&y)

	for i=0;i<m;i++ {
		x = i
		l = i
		for j=1;j<n;j++ {
			l=l*x%m
		}

		if (l%m==y) {
			if f>0 {
				fmt.Print(" ")
			}
			fmt.Print(x)
			f++
		}
	}
	if f==0 {
		fmt.Print("-1")
	}
	fmt.Print("\n")
}