package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n, m, aa, bb int
    var a, b []int
    
    fmt.Scan(&n)
    fmt.Scan(&m)
    
    for i := 0; i <n; i++ {
        fmt.Scan(&aa)
        a = append(a, aa)
    }
    for i := 0; i < m; i++ {
        fmt.Scan(&bb)
        b = append(b, bb)
    }
    
    var c int
    f := lcmList(a)
    g := gcdList(b)
    for i:=f; i <= g;i+=f {
        if g%i == 0{
            c++
        }
    }
    fmt.Println(c)
}

func gcdList(arr []int) int {
    r := arr[0]
    for _, a := range arr {
        r = gcd(r,a)
    }
    return r
}

func lcmList(arr []int) int {
    r := arr[0]
    for _, a := range arr {
        r = lcm(r, a)
    }
    return r
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func lcm(a, b int) int {
    return a * b /gcd(a,b)
}
