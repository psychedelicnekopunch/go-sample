
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


func test(i int) bool {
	switch i {
	case 1:
		return true
	}
	return false
}


func main() {
	checkStr("a")// -- this is A. end. --
	checkStr("test")// -- default. --
	checkStr("b")// -- this is B. end. --

	checkInt(1)// -- i < 5 end. --
	checkInt(4)// -- i < 5 end. --
	checkInt(7)// -- i < 10 end. --

	fmt.Print(test(1), "\n")
	fmt.Print(test(2), "\n")

	for i := 0; i < 10; i++ {
		switch i {
		case 2:
			fmt.Print(i, ": 2\n")
			// break
		default:
			fmt.Print(i, ": その他\n")
			continue
		}
		break
	}
}
