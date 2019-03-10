package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	var n, x, y, fx, fy, i, j int64
	var s1, s2, s3 string
	var cpx, px Pixel
	var pixels [12][12]bool
	ch2 := make(chan Pixel, 100)
	ch := make(chan Pixel, 100)
	firstMethod := false

	fmt.Scan(&s1)
	fmt.Scan(&s2)
	fmt.Scan(&s3)

	isInt, _ := strconv.ParseInt(s3, 10, 64)
	if isInt>0 {
		firstMethod=true
	}

	if firstMethod {
		n, _ = strconv.ParseInt(s1, 10, 64)
		fx, _ = strconv.ParseInt(s2, 10, 64)
		fy, _ = strconv.ParseInt(s3, 10, 64)
		pixels[fx][fy]=true

	} else {
		fx, _ = strconv.ParseInt(s1, 10, 64)
		fy, _ = strconv.ParseInt(s2, 10, 64)
		cpx = Pixel{fx,fy, s3}
		ch <- cpx
		s1 = s3
	}

	if firstMethod {
		for i = 1; i < n; i++ {
			fmt.Scan(&x)
			fmt.Scan(&y)
			pixels[x][y]=true
		}

		fmt.Println(fx, fy)

		var used [12][12]bool
		used[fx][fy] = true


		p := Pixel{fx,fy, ""}
		ch <- p

		do := true
		for do {
			finStr := ""
			p = <- ch

			//right
			x = p.x+1
			y = p.y
			tryAdd(x, y, pixels, &used, ch2, &finStr, "R")

			//top
			x = p.x
			y = p.y+1
			tryAdd(x, y, pixels, &used, ch2, &finStr, "T")

			//left
			if x>0 {
				x = p.x-1
				y = p.y
				tryAdd(x, y, pixels, &used, ch2, &finStr, "T")
			}

			//bottom
			if y>0 {
				x = p.x
				y = p.y-1
				tryAdd(x, y, pixels, &used, ch2, &finStr, "B")
			}

			if len(ch) == 0 {
				for len(ch2)>0 {
					p = <- ch2
					ch <- p
				}
			}
			if len(ch) == 0 {
				do = false
				finStr+="."
			} else {
				finStr+=","
			}
			fmt.Println(finStr)
		}
	} else {
		//считываем все строки и создаем матрицу
		n=0
		do := true
		for do {
			cpx = <- ch
			pixels[cpx.x][cpx.y]=true
			n++

			spl := strings.Split(cpx.next, "")
			for _, c := range spl {
				switch c {
				case "R":
					fmt.Scan(&s1)
					px = Pixel{cpx.x+1,cpx.y, s1}
					ch <- px
				case "T":
					fmt.Scan(&s1)
					px = Pixel{cpx.x,cpx.y+1, s1}
					ch <- px
				case "L":
					fmt.Scan(&s1)
					px = Pixel{cpx.x-1,cpx.y, s1}
					ch <- px
				case "B":
					fmt.Scan(&s1)
					px = Pixel{cpx.x,cpx.y-1, s1}
					ch <- px
				}
			}
			if len(ch) == 0 {
				do = false
			}
		}

		//проходимся по матрице и выводим все пиксели
		fmt.Println(n)
		for i=1;i<12;i++ {
			for j=1;j<12;j++ {
				if pixels[i][j] {
					fmt.Println(i, j)
				}
			}
		}
	}
}

func tryAdd(x int64, y int64, pixels [12][12]bool, used *[12][12]bool, ch2 chan Pixel, finStr *string, outp string) {
	if (pixels[x][y]==true && used[x][y]==false) {
		used[x][y]=true
		*finStr+=outp
		p := Pixel{x,y,""}
		ch2 <- p
	}
}

type Pixel struct {
	x int64;
	y int64;
	next string;
}