
package main


import (
	"fmt"
)


type Test struct {
	ID int
}


func main() {
	var test *Test
	fmt.Print(test)// nil

	test = new(Test)
	fmt.Print(test)// &{0}
}