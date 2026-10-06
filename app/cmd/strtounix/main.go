
package main


import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	s := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Printf("%T, %v\n", s, s)

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Print("error: ", err.Error())
		return
	}
	fmt.Printf("%T, %v\n", i, i)
}
