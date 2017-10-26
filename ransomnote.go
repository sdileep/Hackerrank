package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var m, n int
    var a, b string
    fmt.Scan(&m)
    fmt.Scan(&n)
    
    var mw []string
    var rw []string
    for i := 0; i < m; i++ {
        fmt.Scan(&a)
        mw = append(mw, a)
    }

    for i := 0; i < n; i++ {
        fmt.Scan(&b)
        rw = append(rw, b)
    }
    
    if(verify(mw, rw)) {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }
}

func verify(s1, s2 []string) bool {
    m := mapify(s1)
    for _, e := range s2 {
        if _, ok := m[e]; ok {
            m[e]--
            if m[e] == 0 {
                delete(m, e)
            }
        } else {
            return false
        }
    }
    return true
}

func mapify(s []string) map[string]int {
    m := make(map[string]int)
    for _, e := range s {
        if _, ok := m[e]; ok {
            m[e]++
        } else {
            m[e] = 1
        }
    }
    return m
}
