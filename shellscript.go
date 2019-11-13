
package main


import (
	"fmt"
	"os/exec"
)


func main() {
	output, err := exec.Command("./sh/test").Output()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(string(output))
}