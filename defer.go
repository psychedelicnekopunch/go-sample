package main

import (
	"fmt"
)


func main() {
	fmt.Printf("main 1\n")
	test()
	fmt.Printf("main 2\n")
}


func test() {
	fmt.Printf("test 1\n")
	defer fmt.Printf("test 2\n")

	fmt.Printf("test 3\n")
}