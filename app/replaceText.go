
package main


import (
	"fmt"
	"strings"
)


func main() {
	str := "11111,222222,333333,44444,555555,66666,7777777,8888,99,1000"
	fmt.Print(strings.Replace(str, ",", "，", -1))
}
