
package main


import (
	"fmt"
	"time"
)


func main() {

	i := 0

	fmt.Print("start\n")
	recursion(i)
	fmt.Print("finish\n")
}


func recursion(i int) {
	fmt.Print(i, "\n")
	if i >= 10 {
		return
	}
	i++
	time.Sleep(time.Second * 1)
	recursion(i)
}
