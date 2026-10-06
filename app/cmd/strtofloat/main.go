package main


import (
	"fmt"
	"strconv"
)


func main() {
	parse("1.2345678")
	parse("E1.2345678")
	parse("float64")
	parse("1E6")
	parse("E6")
	parse("1E")
}


func parse(v string) {
	f, err := strconv.ParseFloat(v, 64)
	fmt.Print(f, " : ")
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print("\n")
}
