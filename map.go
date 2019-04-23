
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
	Name string
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


	IDs := map[int]*User{}

	if IDs[10] == nil {
		IDs[10] = &User{
			ID: 10,
			Name: "test",
		}
	}

	print(IDs)
	print(IDs[10].Name)
}