package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n)
	fmt.Scan(&p)

	fmt.Println(min(fromBack(n, p), fromFront(n, p)))
}

func fromFront(a, b int) int {
	return b / 2
}

func fromBack(a, b int) int {
	if a%2 == 0 {
		a++
	}
	return (a - b) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
