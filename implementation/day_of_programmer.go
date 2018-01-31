package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var y int
    fmt.Scan(&y)
    
    if(y == 1918) {
        fmt.Println("26.09.1918")
    } else if(y < 1918 && isLeapYearJulian(y) || (y > 1918 && isLeapYearGregorian(y))) {
        fmt.Printf("12.09.%d\n", y)
    } else {
        fmt.Printf("13.09.%d\n", y)
    }
}

func isLeapYearGregorian(y int) bool {
    return (y % 400 == 0 || ( y % 4 == 0 && y % 100 != 0))
}

func isLeapYearJulian(y int) bool {
    return y % 4 == 0
}
