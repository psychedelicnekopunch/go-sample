
package main


import (
	"fmt"
	"time"
)


func main() {
	s := float64(time.Now().UnixNano())
	var nano float64 = 1000000000
	for i := 0; i < 1000000; i++ {
		for i2 := 0; i2 < 10000; i2++ {}
	}
	f := float64(time.Now().UnixNano())
	fmt.Printf("%.3fsec", (f - s) / nano)
}
