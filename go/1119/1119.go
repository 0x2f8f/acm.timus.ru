package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	var x, y, n, k, l, i, j int
	var d float64
	d = math.Sqrt(20000)
	fmt.Scan(&x)
	fmt.Scan(&y)

	B, C := initMatrix(x, y)

	fmt.Scan(&n)
	for ; n > 0; n-- {
		fmt.Scan(&k, &l)
		i = getInd((k - 1), (l - 1), B)
		j = getInd(k, l, B)
		C[i][j] = d
	}

	//Floyd–Warshall algorithm
	max := x * y * 2
	for k = 0; k < max; k++ {
		for i = 0; i < max; i++ {
			for j = 0; j < max; j++ {
				C[i][j] = math.Min(C[i][j], C[i][k]+C[k][j])
			}
		}
	}
	fmt.Println(Round(C[0][(max-1)]))
}

func initMatrix(x, y int) (map[string]int, [][]float64) {
	var max, k int
	max = x * y * 2

	A := make(map[int]string)
	B := make(map[string]int)
	C := make([][]float64, max)

	for i := range C {
		C[i] = make([]float64, max)
	}

	n := 0
	var ind string
	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			ind = strconv.Itoa(j) + "." + strconv.Itoa(i)
			A[n] = ind
			B[ind] = n

			if (j > 0) {
				k = getInd(j-1, i, B)
				C[k][n] = 100
			}

			if (i > 0) {
				k = getInd(j, i-1, B)
				C[k][n] = 100
			}

			n++
		}
	}

	for i := 0; i < max; i++ {
		for j := 0; j < max; j++ {
			if (i != j) {
				if (C[i][j] == 0) {
					C[i][j] = 2000000000
				}
			}
		}
	}

	return B, C
}

func getInd(x, y int, B map[string]int) int {
	return B[strconv.Itoa(x)+"."+strconv.Itoa(y)]
}

func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}