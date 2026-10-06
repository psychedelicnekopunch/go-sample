
package main


import (
	"fmt"
)


func main() {
	s := "sample"
	b := []byte(s)
	fmt.Print(b, "\n")
	fmt.Print(string(b), "\n")
}
