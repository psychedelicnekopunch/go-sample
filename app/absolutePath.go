
package main


import (
	"fmt"
	"path/filepath"
)


func main() {
	path, err := filepath.Abs("./")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	fmt.Print(path)
}
