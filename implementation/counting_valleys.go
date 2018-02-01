func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	var b string
	fmt.Scan(&n)
	fmt.Scan(&b)

	fmt.Println(calculateValley(b))
}

func calculateValley(a string) int {
	var d, v int
	for _, runeValue := range a {
		if runeValue == 85 {
			d++
			if d == 0 {
				v++
			}
		} else {
			d--
		}
	}
	return v
}
