package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n, a int
    var arr []int
    fmt.Scan(&n)
    
    for i:=0;i<n;i++ {
        fmt.Scan(&a)
        arr = append(arr, a)
    }
    mc := BirthdayCakeCandles(arr)
    fmt.Println(mc)
}

func BirthdayCakeCandles(arr []int) int {
	m := arr[0]
	mc := 1
	for i:=1;i<len(arr);i++{
			if arr[i] > m {
					mc = 0
					m = arr[i]
			}
			if arr[i] == m {
					mc++
			}
	}
	return mc
}
