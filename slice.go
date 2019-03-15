
package main


import (
	"fmt"
)

type User struct {
	ID int
	Name string
}



func print(v interface{}) {
	fmt.Print(v)
	fmt.Print("\n")
}



func main() {
	// その1
	var sample []User
	print(sample)

	sample = append(sample, User{
		ID: 1,
		Name: "test",
	})
	print(sample)

	// その2
	sample2 := []User{}
	print(sample2)

	sample2 = append(sample2, User{
		ID: 2,
		Name: "test2",
	})
	print(sample2)
}