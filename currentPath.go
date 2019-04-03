package main

import (
	"os"
	"fmt"
)

func main() {
	currentPath, _ := os.Getwd()
	fmt.Print(currentPath)
}