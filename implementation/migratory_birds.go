package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, a int
	fmt.Scan(&n)

	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[arr[i]]++
	}

	var max, maxPos int
	for k, v := range m {
		if v >= max {
			max = v
			maxPos = k
		}
	}
	fmt.Println(maxPos)
}
