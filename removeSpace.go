package main


import (
	"fmt"
	"strings"
)


func main() {

	value := " a b c d e f g "
	s := strings.Replace(value, " ", "", -1)

	fmt.Printf("[%s]\n", value)
	fmt.Printf("[%s]\n", s)
}
