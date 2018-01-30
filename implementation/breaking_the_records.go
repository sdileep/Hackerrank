package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n,s int
    fmt.Scan(&n)
    var arr []int
    for i := 0; i < n; i++ {
        fmt.Scan(&s)
        arr = append(arr, s)
    }
    hs, ls := calculateHighLowest(arr)
    fmt.Printf("%v %v\n", hs, ls)
}

func calculateHighLowest(arr []int) (int, int) {
    var hc, lc int
    hs, ls := arr[0], arr[0]
    for _, s := range arr {
        if s > hs {
            hs = s
            hc++
        } else if s < ls {
            ls = s
            lc++
        }
    }
    return hc, lc
}
