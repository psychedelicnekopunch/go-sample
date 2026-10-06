package main

import (
	"fmt"
)


func main() {

	var strA *string
	var strB *string

	// TEST 1
	if strA == strB {
		fmt.Print("TEST 1: Equal\n")
	} else {
		fmt.Print("TEST 1: Not Equal\n")
	}
	fmt.Print(strA, "\n")
	fmt.Print(strB, "\n")


	// TEST 2
	str := "test"
	str2 := "test"

	strA = &str
	strB = &str2

	if strA == strB {
		fmt.Print("TEST 2: Equal\n")
	} else {
		fmt.Print("TEST 2: Not Equal\n")
	}
	fmt.Print(strA, "\n")
	fmt.Print(strB, "\n")


	// TEST 3
	if *strA == *strB {
		fmt.Print("TEST 3: Equal\n")
	} else {
		fmt.Print("TEST 3: Not Equal\n")
	}
	fmt.Print(*strA, "\n")
	fmt.Print(*strB, "\n")
}
