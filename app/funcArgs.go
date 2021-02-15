package main

import (
	"fmt"
	"strconv"
)


func testFunc(value string, values ...string) {
	fmt.Print(value + " ======== \n")
	fmt.Print(values)
	for i, value := range values {
		if i == 0 { fmt.Print("\n") }
		fmt.Print(strconv.Itoa(i) + ":" + value + " ")
	}
	fmt.Print("\n\n")
}


func testFunc2(value string, values ...string) {
	testFunc(value, values...)
}


func main() {
	testFunc("test1")
	testFunc("test2", "test2-a")
	testFunc("test3", "test3-a", "test3-b", "test3-c",)


	testFunc2("test4")
	testFunc2("test5", "test5-a")
	testFunc2("test6", "test6-a", "test6-b", "test6-c",)
}
