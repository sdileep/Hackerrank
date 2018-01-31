package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n, k, c, b int
    fmt.Scan(&n)
    fmt.Scan(&k)
    var arr []int
    
    for i := 0; i < n; i++ {
        fmt.Scan(&c)
        arr = append(arr, c)
    }
    fmt.Scan(&b)
    
    count := 0
    for i := 0; i < n; i++ {
        if i != k {
            count+=arr[i]
        }
    }
    charge := count / 2
    if charge == b {
        fmt.Println("Bon Appetit")
    } else {
        fmt.Println(b - charge)
    }
}
