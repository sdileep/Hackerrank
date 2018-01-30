package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, a, d, m int
	fmt.Scan(&n)
	var c []int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		c = append(c, a)
	}
	fmt.Scan(&d)
	fmt.Scan(&m)

	sum := 0
	total := 0
	for i := 0; i < n-m+1; i++ {
		for j := i; j < i+m; j++ {
			sum += c[j]
		}
		if sum == d {
			total++
		}
		sum = 0
		// fmt.Println("comparing ..", c[i], " with ", c[i+1])
	}
	fmt.Println(total)
}
