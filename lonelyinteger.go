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
    fmt.Println(uniqueNum(arr))
}

func uniqueNum(arr []int) int {
    var res int
    for i := 0; i < len(arr);i++ {
        res^= arr[i] 
    }
    return res
}
