
package main


import (
	"fmt"
)


func main() {
	s := [3]int{}
	s[1] = 2
	change(s)
	fmt.Print("s is " , s, "\n")
}


func change(v [3]int) {
	v[0] = 10
}
