package main


import (
	"fmt"
	"strings"
)


func main() {

	str := "2019.09.13,2019.09.14,2019.09.15,2019.09.16,2019.09.17,2019.09.18,2019.09.19,2019.09.20"

	arr := []string{}
	arr = strings.Split(str, ",")

	for _, v := range arr {
		fmt.Print(v + "\n")
	}


	str2 := "127.0.0.1:45634"
	arr2 := []string{}
	arr2 = strings.Split(str2, ":")

	for i, v := range arr2 {
		fmt.Print(i, ": ", v + "\n")
	}
}
