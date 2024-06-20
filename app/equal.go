
package main


import (
	"fmt"
	// "testing"
	// "github.com/stretchr/testify/assert"
)


type Test struct {
	ID int
}



func main() {

	// t := *testing.T
	// assert.Equal(t, 123, 123, "they should be equal")

	var i int
	i2 := 0

	fmt.Print("i == i2: ", i == i2, "\n")


	var t Test
	t2 := Test{}
	t3 := Test{ ID: 1 }

	fmt.Print("t == t2: ", t == t2, "\n")
	fmt.Print("t == t3: ", t == t3, "\n")

}
