
package main


import (
	"fmt"
)


// type Color int
type Color string


const (
	// Red Color = iota
	// Blue
	// Yellow
	Red Color = "red"
	Blue Color = "blue"
	Yellow Color = "yellow"
)


func main() {

	// var red Color = Blue
	var red Color = "black"
	fmt.Print(red, "\n")

}
