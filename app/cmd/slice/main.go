
package main


import (
	"fmt"
)

type User struct {
	ID int
	Name string
}


func main() {
	// その1
	var sample []User
	fmt.Print(sample, "\n")

	sample = append(sample, User{
		ID: 1,
		Name: "test",
	})
	fmt.Print(sample, "\n")

	// その2
	sample2 := []User{}
	fmt.Print(sample2, "\n")

	sample2 = append(sample2, User{
		ID: 2,
		Name: "test2",
	})
	fmt.Print(sample2, "\n")


	strs := []string{
		"a",
		"b",
		"c",
	}
	fmt.Print(strs, "\n")
}
