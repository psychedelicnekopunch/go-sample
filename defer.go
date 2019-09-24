package main

import (
	"fmt"
)


func main() {
	fmt.Printf("main: 1\n")
	test()
	fmt.Printf("main: 2\n")
	test2()
}


func test() {
	fmt.Printf("test: 1\n")

	defer fmt.Printf("test: 2\n")

	fmt.Printf("test: 3\n")
}


func test2() {
	fmt.Printf("test2: 1\n")

	defer fmt.Printf("test2: 2\n")
	defer fmt.Printf("test2: 3\n")
	defer fmt.Printf("test2: 4\n")

	fmt.Printf("test2: 5\n")
}
