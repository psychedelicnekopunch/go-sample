package main


import (
	"fmt"
	"strings"
)


func main() {
	convert(0)
	convert(10)
	convert(230)
	convert(4500)
	convert(67000)
	convert(890000)
	convert(1200000)
	convert(34000000)
}


func convert(integer int) {
	str := fmt.Sprintf("%d", integer)
	arr := strings.Split(str, "")
	cnt := len(arr) - 1
	// fmt.Print(arr)
	res := ""
	i2 := 0
	for i := cnt; i >= 0; i-- {
		str = fmt.Sprintf("0%s", str)
		// fmt.Print(i2 % 3)
		if i2 > 2 && i2 % 3 == 0 {
			res = fmt.Sprintf(",%s", res)
		}
		res = fmt.Sprintf("%s%s", arr[i], res)
		i2++
	}
	// fmt.Print("\n" + res + "\n\n")
	fmt.Print(res + "\n")
}