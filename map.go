
package main


import (
	"fmt"
)


func print(v interface{}) {
	fmt.Print(v)
	fmt.Print("\n")
}


type User struct {
	ID int
}


func main() {
	// その1
	var sample map[string]interface{}
	print(sample)
	sample = map[string]interface{}{}
	sample["ID"] = 1
	sample["Name"] = "test"
	print(sample)

	// その2
	sample2 := map[string]interface{}{}
	print(sample2)
	sample2["ID"] = 2
	sample2["Name"] = "test2"
	print(sample2)
}