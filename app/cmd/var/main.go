
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

	var test2 Test
	fmt.Print(test2)// {0}

	var test3 int
	test3 = 0

	var _ int
	_ = 0

	test4 := *Test{}

	var test5 *int
	test5 = test3
}
