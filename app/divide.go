
package main

import (
	"fmt"
)


func main() {
	fmt.Print("0 / 10 = ", 0 / 10, "\n")
	// fmt.Print("10 / 0 = ", 10 / 0, "\n")// invalid operation: division by zero
	fmt.Print("10 / 0 = ", float64(10) / getZeroAsFloat64(), "\n")
	fmt.Print("10 / 0 = ", 10 / getZeroAsInt(), "\n")// panic: runtime error: integer divide by zero
}


func getZeroAsFloat64() float64 {
	return 0
}

func getZeroAsInt() int {
	return 0
}
