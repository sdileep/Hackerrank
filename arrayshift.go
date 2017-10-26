package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n, k, a int
    fmt.Scan(&n)
    fmt.Scan(&k)
    var arr []int
    for i:=0; i < n; i++ {
        fmt.Scan(&a)
        arr = append(arr, a)
    }
    arrayShift(arr, k)
    for _, e := range arr {
        fmt.Printf("%d ", e)
    }
    fmt.Println()
}

func arrayShift(nums []int, k int) {
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	reverse(nums, 0, len(nums)-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
