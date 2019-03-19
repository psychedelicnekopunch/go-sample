
package main


import (
	"fmt"
	"strconv"
)


func main() {

	lists := []string{ "a","b","c","d","e","f","g","h","i" }
	cnt := len(lists)

	// for
	for i := 0; i < cnt; i++ {
		// 0:a 1:b 2:c 3:d 4:e 5:f 6:g 7:h 8:i
		fmt.Printf(strconv.Itoa(i) + ":" + lists[i] + " ")
	}
	fmt.Printf("\n")


	// foreach
	for key, value := range lists {
		// 0:a 1:b 2:c 3:d 4:e 5:f 6:g 7:h 8:i
		fmt.Printf(strconv.Itoa(key) + ":" + value + " ")
	}
	fmt.Printf("\n")


	// while
	i := 0
	for i < cnt {
		// 0:a 1:b 2:c 3:d 4:e 5:f 6:g 7:h 8:i
		fmt.Printf(strconv.Itoa(i) + ":" + lists[i] + " ")
		i++
	}
	fmt.Printf("\n")
}