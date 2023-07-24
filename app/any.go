
package main


import (
	"fmt"
)

func main() {

	// var test any = "test"
	var test any = true
	// var test interface{} = "test"
	// fmt.Print(test, "\n")
	// t, valid := test.(string)
	t, valid := test.(bool)
	// t, valid := test.(int)
	fmt.Print(t, ": ", valid, "\n")

	switch test.(type) {
	case string:
		fmt.Print("this type is String.\n")
	case int:
		fmt.Print("this type is Int.\n")
	case bool:
		fmt.Print("this type is Boolean.\n")
	}
}
