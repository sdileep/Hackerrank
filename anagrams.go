package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var s1, s2 string
    fmt.Scan(&s1)
    fmt.Scan(&s2)
    fmt.Println(compare(s1, s2))
}

func compare(s1, s2 string) int {
	m1 := mapify(s1)
	m2 := mapify(s2)

	return countOccurreces(m1, m2) + countOccurreces(m2, m1)
}

func countOccurreces(m1, m2 map[string]int) int {
	var count int
	for k, v := range m1 {
		if _, ok := m2[k]; ok {
			count += abs(v - m2[k])
			delete(m2, k) // no need to process same
		} else {
			count += v
		}
	}
	return count
}

func mapify(s string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		st := string(s[i])
		if _, ok := m[st]; ok {
			t := m[st]
			t++
			m[st] = t
		} else {
			m[st] = 1
		}
	}
	return m
}

func abs(a int) int {
    if a < 0 {
        return a * -1
    }
    return a
}
