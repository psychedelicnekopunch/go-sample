
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


	// var IDs map[int]*User // panic: assignment to entry in nil map
	IDs := map[int]*User{}
	print(IDs)

	if IDs[10] == nil {
		IDs[10] = &User{
			ID: 10,
			Name: "test",
		}
	}

	print(IDs)
	print(IDs[10])
	print(IDs[10].Name)

	Users := map[int]User{}
	print(Users)

	Users[1] = User{
		ID: 1,
		Name: "test2",
	}

	print(Users)
	print(Users[1])
	print(Users[1].Name)
	print(Users[10])
	print(Users[10].Name)


	Nums := map[int]int{}
	print(Nums[0])


	// panic: assignment to entry in nil map
	var test map[int]string
	// test = map[int]string{}
	// test := map[int]string{}
	test[0] = "test"
}