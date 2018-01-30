package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, k, a int
	fmt.Scan(&n)
	fmt.Scan(&k)

	var arr []int
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	count := 0
	buckets := make([]int, k)
	for i := 0; i < n; i++ {
		bucket := arr[i] % k
		count += buckets[(k-bucket)%k]
		buckets[bucket]++
	}
	fmt.Println(count)
}
