
package main


import (
	"fmt"
)


type Person struct {
	ID int
	Name struct {
		First string
		Last string
	}
}


type Person2 struct {
	ID int
	Name Name
}


type Name struct {
	First string
	Last string
}


func main() {
	p := Person{
		ID: 1,
		Name: struct{
			First string
			Last string
		}{
			First: "neco",
			Last: "nekomura",
		},
	}
	fmt.Print(p)


	p2 := Person2{
		ID: 2,
		Name: Name{
			First: "bar",
			Last: "murata",
		},
	}
	fmt.Print(p2)
}
