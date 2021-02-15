package main


import (
	"fmt"
	"math"
)


func main() {
	var i int = 123
	var i2 int = 3456

	f := float64(i) / float64(i2)
	fmt.Print(f)
	fmt.Print("\n")
	fmt.Print(math.Floor(f * 100))
	fmt.Print("\n")

	fmt.Printf("%.1f\n", f)
	fmt.Printf("%.3f\n", f)
}
