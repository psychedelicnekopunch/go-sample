
package main


import (
	"fmt"
)


func main() {
	i := 1
	add(i)
	fmt.Print(i, "\n")

	add2(&i)
	fmt.Print(i, "\n")
}


func add(v int) {
	v += 1
}


func add2(v *int) {
	*v += 1
}
