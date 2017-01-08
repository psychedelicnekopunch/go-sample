package main

import (
	"fmt"
	"go-sample/sample"
	sampleTest "go-sample/sample"
	. "go-sample/sample"
	"go-sample/sample2"
)

func FuncSample() {
	fmt.Printf("func sample in main\n")
}

func main() {
	sample.Func()
	sampleTest.Func()
	Func2()
	sample2.Func3()
	sample2.FuncSample()
	FuncSample()
}