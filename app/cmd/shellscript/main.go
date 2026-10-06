
package main


import (
	"fmt"
	"os/exec"
)


func main() {
	output, err := exec.Command("scripts/shellscript/test").Output()
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}
	fmt.Print(string(output))
}
