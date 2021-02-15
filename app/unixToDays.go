
package main


import (
	"fmt"
	"time"
)

func main() {
	unix := time.Now().Unix()
	// unix2 := time.Now().Unix() + (60*60*24*10)
	unix2 := time.Now().Unix() + (60*60*24*10) + 123792
	// unix2 := time.Now().Unix() + 123792
	s := unix2 - unix
	fmt.Print(s, "\n")
	fmt.Print(s / (60*60*24), "\n")
}
