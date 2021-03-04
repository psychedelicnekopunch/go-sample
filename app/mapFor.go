
package main


import (
	"fmt"
)


func main() {

	mapInts := map[int]int{}

	for i := 0; i < 10; i++ {
		mapInts[i + 1] = i + 1
	}
	for i, v := range mapInts {
		fmt.Print(i, " : ", v, "\n")
	}


	mapStrs := map[string]int{}
	mapStrs["a"] = 1
	mapStrs["b"] = 2
	mapStrs["c"] = 3
	mapStrs["d"] = 4
	mapStrs["e"] = 5

	for i, v := range mapStrs {
		fmt.Print(i, " : ", v, "\n")
	}
}
