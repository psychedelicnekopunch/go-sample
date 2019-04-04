package main

import (
	"fmt"
	"psychedelicnekopunch/go-sample/sample"
	sampleTest "psychedelicnekopunch/go-sample/sample"
	. "psychedelicnekopunch/go-sample/sample"
	"psychedelicnekopunch/go-sample/sample2"
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