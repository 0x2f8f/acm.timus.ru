package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

var gnoms int64 = 7

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	var n int64
	ch := make(chan int64, 100)

	var s string
	fmt.Scan(&s)
	spl := strings.Split(s, "")
	for _, c := range spl {
		n, _ = strconv.ParseInt(c, 10, 64)
		ch <- n
	}

	n=0
	do := true
	for do {
		k := <- ch

		n = n * 10 + k
		if n >= gnoms {
			n=n%gnoms
		}

		if len(ch) == 0 {
			do = false
		}
	}
	fmt.Println(n)
}