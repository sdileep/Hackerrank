package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var a, t int
    var arr []int
    for i := 0;i<5;i++ {
        fmt.Scan(&a)
        arr = append(arr, a)
    }
    min := arr[0]
    max := arr[0]
    for i :=0;i<len(arr);i++ {
        t+=arr[i]
        if arr[i] < min {
            min = arr[i]
        }
        if arr[i] > max {
            max = arr[i]
        }
    }
    fmt.Printf("%d %d\n", t-max, t-min)
}
