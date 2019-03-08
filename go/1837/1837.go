package main

import (
	"fmt"
	"os"
)

var isN int //порядковый номер Исенбаева

func main() {
	_, err := os.Stat("in.txt")
	if err == nil {
		os.Stdin, err = os.OpenFile("in.txt", os.O_RDONLY, 0666)
	}

	m := make(map[string]int) //асс-ый массив, где ключ - фамилия, значение - порядковый номер

	var n int //кол-во команд

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
/*
	for key, value := range m {
		fmt.Println(key, " ", value)
	}

	fmt.Println(graph)
*/
	fmt.Println(isN)
}

//добавление фамилии в справочник
func addItemMap(m map[string]int, f string) {
	_, isset := m[f]
	if !isset {
		m[f] = len(m)
		if f == "Isenbaev" {
			isN = m[f]
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