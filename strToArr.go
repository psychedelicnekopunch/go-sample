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
}