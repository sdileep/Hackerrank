package main

import "fmt"

func main() {
	var n, a int
	var m [][]int
	
	fmt.Scan(&n)
	for i :=0; i < n; i++ {
		var r []int
		for j := 0; j <n;j++ {
				fmt.Scan(&a)
				r = append(r, a)
		}
		m = append(m, r)
	}
	dd := DiagonalDifference(m)
	fmt.Println(dd)
}

func DiagonalDifference(m [][]int) int {
	var p,s int
	n := len(m)
	for i :=0;i<n;i++ {
			p += m[i][i]
			s += m[i][n-1-i]
	}
	return abs(p-s)
}

func abs(x int) int {
	if x < 0 {
		x *=-1
	}
	return x
}

