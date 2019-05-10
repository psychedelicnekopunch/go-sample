
package main


import (
	"fmt"
)


func checkStr(str string) {
	switch str {
	case "a":
		fmt.Print("-- this is A.")
	case "b":
		fmt.Print("-- this is B.")
	default:
		fmt.Print("-- default. --\n")
		return
	}
	fmt.Print(" end. --\n")
}


func checkInt(i int) {
	switch {
	case i < 5:
		fmt.Print("-- i < 5")
	case i < 2:
		fmt.Print("-- i < 2")
	case i < 10:
		fmt.Print("-- i < 10")
	default:
		fmt.Print("-- default. --\n")
		return
	}
	fmt.Print(" end. --\n")
}


func main() {
	checkStr("a")// -- this is A. end. --
	checkStr("test")// -- default. --
	checkStr("b")// -- this is B. end. --

	checkInt(1)// -- i < 5 end. --
	checkInt(4)// -- i < 5 end. --
	checkInt(7)// -- i < 10 end. --
}
