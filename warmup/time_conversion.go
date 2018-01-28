package main
import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var t string
    fmt.Scan(&t)
    
    time := strings.Split(t, ":")
    if strings.HasPrefix(t, "12") && strings.HasSuffix(t, "AM") {
        // replace 12 with 00
        fmt.Printf("%s:%s:%s\n","00",time[1],time[2][0:2])
    } else if strings.HasSuffix(t, "AM") || strings.HasPrefix(t, "12") && strings.HasSuffix(t, "PM"){
			fmt.Printf("%s:%s:%s\n",time[0],time[1],time[2][0:2])
		} else  {
			i, _ := strconv.Atoi(time[0])
			fmt.Printf("%v:%s:%s\n",i+12,time[1],time[2][0:2])
		}
}
