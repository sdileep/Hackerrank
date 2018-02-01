package main

import "fmt"

func main() {
	var n, c int
	fmt.Scan(&n)

	var arr []int

	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		arr = append(arr, c)
	}
	fmt.Println(getPairs(arr))
}

func getPairs(arr []int) int {
	m := make(map[int]int)
	var t int
	for _, a := range arr {
		if _, ok := m[a]; ok {
			t++
			delete(m, a)
		} else {
			m[a] = 1
		}
	}
	return t
}
