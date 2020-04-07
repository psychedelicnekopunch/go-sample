package main


import (
	"fmt"
	"strconv"
)


func main() {
	var i string = "1.2345678"

	f, err := strconv.ParseFloat(i, 64)

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Print(f)
	// fmt.Printf("%.3f\n", f)
}
