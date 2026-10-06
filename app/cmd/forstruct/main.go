
package main


import (
	"fmt"
	"reflect"
)


type Test struct {
	A string
	B bool
	C int
	D string
}


func NewTest() *Test {
	return &Test{
		A: "this is A.",
		B: true,
		C: 99,
		D: "this is D.",
	}
}

func main() {
	test := NewTest()
	t := reflect.TypeOf(*test)
	elem := reflect.ValueOf(test).Elem()
	cnt := elem.NumField()
	// cnt := t.NumField()
	for i := 0; i < cnt; i++ {
		fmt.Print(t.Field(i).Name)
		fmt.Print(": ")
		fmt.Print(elem.Field(i))
		// fmt.Print(elem.Field(i).Interface())
		fmt.Print("\n")
	}

	// cannot range over test (type *Test)
	// for _, t := range test {
	// 	fmt.Print(t)
	// }
}
