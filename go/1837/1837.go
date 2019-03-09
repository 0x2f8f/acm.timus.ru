package main

import (
	"fmt"
	"os"
	"sort"
)

var isenbaevNum int //порядковый номер Исенбаева

func main() {
	isenbaevNum = -1
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	m := make(map[string]int) //асс-ый массив, где ключ - фамилия, значение - порядковый номер

	var n int //кол-во команд
	var l int //временная переменная для приема значения из канала

	var a, b, c string
	fmt.Scan(&n)

	nm := n*3 //max возможное кол-во людей
	graph := make([][]bool, nm, nm) //матрица графа

	for i:=0; i<nm; i++ {
		graph[i] = make([]bool, nm, nm)
	}

	for i:=0; i<n; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&c)

		//добавляем в ассоциативный массив
		addItemMap(m, a)
		addItemMap(m, b)
		addItemMap(m, c)

		//добавляем в матрицу графа
		addItemsGraph(graph, m, a, b, c, n)
	}

	//отсортировать map по алфавиту
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, surname := range keys {
		//fmt.Println("Key:", surname, "Value:", m[surname])

		//берем первую вершину и идём до исенбаева
		currentNum := m[surname] //номер текущей вершины
		used := make([]bool, nm, nm) //сгоревшие вершины
		used[currentNum] = true

		i := 0
		val := -1

		//объявляем очередь
		ch := make(chan int, len(m))
		ch2 := make(chan int, len(m))
		ch <- currentNum

		for val==-1 {
			currentNum = <- ch
			//fmt.Println(currentNum, i )
			if currentNum != isenbaevNum {
				addChildren(currentNum, ch2, graph, used)
			} else {
				val = i
			}

			if len(ch) == 0 {

				for len(ch2)>0 {
					l = <- ch2
					ch <- l
				}
				i++
				if len(ch) == 0 && val==-1 {
					val = -2
				}
			}
		}

		if val != -2 {
			fmt.Println(surname, val)
		} else {
			fmt.Println(surname, "undefined")
		}
	}
}

//добавление фамилии в справочник
func addItemMap(m map[string]int, f string) {
	_, isset := m[f]
	if !isset {
		m[f] = len(m)
		if f == "Isenbaev" {
			isenbaevNum = m[f]
		}
	}
}

//добавление команды в матрицу графа
func addItemsGraph(graph [][]bool, m map[string]int, a string, b string, c string, n int)  {
	kA, _ := m[a]
	kB, _ := m[b]
	kC, _ := m[c]

	graph[kA][kB] = true
	graph[kB][kA] = true

	graph[kA][kC] = true
	graph[kC][kA] = true

	graph[kB][kC] = true
	graph[kC][kB] = true
}

//добавление смежных (дочерних) вершин в очередь
func addChildren(current int, ch chan int, graph [][]bool, used []bool){
	for i:=0; i<len(graph[current]); i++ {
		if graph[current][i]==true && used[i]!=true {
			used[i] = true
			ch <- i
		}
	}
}