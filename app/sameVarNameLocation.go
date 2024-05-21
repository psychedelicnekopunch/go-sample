
package main


import (
	"fmt"
)


func main() {
	i := 1

	if true {
		i := 5
		// i = 3
		fmt.Print(i, "\n")
	}

	fmt.Print(i, "\n")


	i, err := getInt(5)

	if err == nil {
		i, _ := getInt(10)
		// i, _ = getInt(10)
		fmt.Print(i, "\n")
	}

	fmt.Print(i, "\n")
}


func getInt(v int) (i int, err error) {
	return v, nil
}
