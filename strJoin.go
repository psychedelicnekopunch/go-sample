
package main


import (
	"fmt"
)

func main() {
	str := "abc"
	str2 := "日本語です。"
	strJoin := str
	strJoin = strJoin + ",\n\n" + str2
	fmt.Printf(strJoin)
}