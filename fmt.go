package main

import (
	"fmt"
)


func main() {
	var f float64

	f = 1.02345
	fmt.Printf(fmt.Sprintf("%f\n", f))
	fmt.Printf(fmt.Sprintf("%g\n", f))

	f = 1
	fmt.Printf(fmt.Sprintf("%f\n", f))
	fmt.Printf(fmt.Sprintf("%g\n", f))
}