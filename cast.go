
package main


import (
	"fmt"
	"strconv"
)


func main() {
	fmt.Printf("[string to bool]\n")
	strToBool("0")// false <nil>
	strToBool("1")// true <nil>
	strToBool("2")// false strconv.ParseBool: parsing "2": invalid syntax
	strToBool("a")// false strconv.ParseBool: parsing "a": invalid syntax
	strToBool("01")// false strconv.ParseBool: parsing "01": invalid syntax
	strToBool("-1")// false strconv.ParseBool: parsing "-1": invalid syntax

	fmt.Printf("[string to int]\n")
	Atoi("0")// 0 <nil>
	Atoi("1.0")// 0 strconv.Atoi: parsing "1.0": invalid syntax
	Atoi("a")// 0 strconv.Atoi: parsing "a": invalid syntax
	Atoi("01")// 1 <nil>
	Atoi("-1")// -1 <nil>
}


func strToBool(str string) {
	fmt.Print(strconv.ParseBool(str))
	fmt.Printf("\n")
}


func Atoi(str string) {
	fmt.Print(strconv.Atoi(str))
	fmt.Printf("\n")
}