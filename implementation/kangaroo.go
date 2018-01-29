package main
import "fmt"

// brute force solution...

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var x1, x2, v1, v2 int
    
    fmt.Scan(&x1)
    fmt.Scan(&v1)
    fmt.Scan(&x2)
    fmt.Scan(&v2)
    
    if doesIntersect(x1, v1, x2, v2) {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

func doesIntersect(x1, v1, x2, v2 int) bool {
	if (x1 > x2 && v1 > v2) || (x2 > x1 && v2 > v1) {
		return false
	}
	for i, j := x1, x2; i < 100000000; i, j = i+v1, j+v2 {
		if i == j {
			return true
		}
	}
	return false
}
