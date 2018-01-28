package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, a int
	var arr []int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}
	scaled := scale(arr)
	for _, s := range scaled {
		fmt.Println(s)
	}
}

func scale(arr []int) []int {
	var r []int
	for i := 0; i < len(arr); i++ {
		if arr[i] < 38 {
			r = append(r, arr[i])
		} else {
			r = append(r, roundUp(arr[i]))
		}
	}
	return r
}

func roundUp(n int) int {
	if n%5 >= 3 {
		n += 5 - (n % 5)
	}
	return n
}
