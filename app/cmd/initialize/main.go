
package main


import (
	"fmt"
)


func main() {

	var str string
	str2 := ""

	fmt.Printf(str)
	fmt.Printf(str2)


	var slice []int
	slice2 := []int{}

	fmt.Println(slice)
	fmt.Println(slice2)

	type Neko struct {
		Name string
	}

	var slice3 []Neko
	slice4 := []Neko{}

	fmt.Println(slice3)
	fmt.Println(slice4)
}
