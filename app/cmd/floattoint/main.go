package main


import (
	"fmt"
	"math"
)


func main() {
	var i float64 = 1.1
	fmt.Printf("%d\n", int(math.Ceil(i)))
	fmt.Printf("%d\n", int(i))
}
