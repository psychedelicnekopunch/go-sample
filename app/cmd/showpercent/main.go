package main


import (
	"fmt"
)


func main() {
	cnt := 12345
	// cnt := 100
	percent := 0
	for i := 0; i <= cnt; i++ {
		newPercent := int(float64(i) / float64(cnt) * 100)
		if percent != newPercent || i == 0 {
			percent = newPercent
			fmt.Print(fmt.Sprintf("%d%s\n", percent, "%"))
		}
	}
}
