package main

import (
	"os"
	"fmt"
)

func main() {
	currentDir, _ := os.Getwd()
	fmt.Print(currentDir)
}
