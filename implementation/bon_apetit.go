package main
import "fmt"

func main() {
 fmt.Scan(&b)
	fs := fairSplit(arr, k)
	if b != fs {
		fmt.Println(b - fs)
	} else {
		fmt.Println("Bon Appetit")
	}
}

func fairSplit(arr []int, k int) int {
	var t int
	for i, c := range arr {
		if i != k {
			t += c
		}
	}
	return t / 2
}
