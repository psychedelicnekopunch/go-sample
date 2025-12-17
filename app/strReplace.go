package main


import (
	// "encoding/json"
	"fmt"
	// "strings"
)


func main() {
	// sample@gmail.com → s*****@gmail.com に変換したい
	//
	str := "sample@gmail.com"
	arr := []rune(str)
	res := ""

	foundAt := false

	for i, v := range arr {
		// fmt.Print(v, "\n")
		// fmt.Print(string(v), "\n")
		s := string(v)

		if i == 0 {
			res = fmt.Sprintf("%s*****", s)
		}
		if s == "@" {
			foundAt = true
		}
		if foundAt {
			res = fmt.Sprintf("%s%s", res, s)
		}
	}
	fmt.Print(res, "\n")
}
