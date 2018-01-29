package main
import "fmt"

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var s, t, a, b, m, n, mm, nn int
    fmt.Scan( & s)
    fmt.Scan( & t)
    fmt.Scan( & a)
    fmt.Scan( & b)
    fmt.Scan( & m)
    fmt.Scan( & n)

    var apples[] int
    var oranges[] int
    for i: = 0;
    i < m;
    i++{
        fmt.Scan( & mm)
        apples = append(apples, mm)
    }

    for i: = 0;
    i < n;
    i++{
        fmt.Scan( & nn)
        oranges = append(oranges, nn)
    }

    fmt.Printf("%d\n%d\n", countFruit(apples, a, s, t), countFruit(oranges, b, s, t))
}

func countFruit(arr[] int, a, s, t int) int {
    var c int
    for _, f: = range arr {
        if a + f >= s && a + f <= t {
            c++
        }
    }
    return c
}
