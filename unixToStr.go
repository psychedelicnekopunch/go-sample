
package main


import (
	"fmt"
	"time"
)

func main() {
	s := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Print(s)
}
