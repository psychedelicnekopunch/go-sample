package main


import (
	"fmt"
	"strings"
)


func main() {
	convert(0)
	convert(30)
	convert(145)
	convert(1150)
}


func convert(integer int) {
	str := fmt.Sprintf("%d", integer)
	cnt := 4 - len(str)
	// fmt.Print(len(str))
	for i := 0; i < cnt; i++ {
		str = fmt.Sprintf("0%s", str)
	}
	// fmt.Print(str + "\n")
	// arr := []string{}
	arr := strings.Split(str, "")
	// fmt.Print(arr)
	res := ""
	for i, v := range arr {
		if i == 2 {
			res = fmt.Sprintf("%s:", res)
		}
		res = fmt.Sprintf("%s%s", res, v)
	}
	fmt.Print(res + "\n")
}
