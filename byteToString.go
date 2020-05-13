
package main


import (
	"fmt"
)


func main() {
	s := "sample"
	b := []byte(s)
	fmt.Print(b)
	fmt.Print(string(b))
}
