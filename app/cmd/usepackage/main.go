package main

import (
	"fmt"
	"github.com/psychedelicnekopunch/go-sample/internal/usepackage"
	usepackageTest "github.com/psychedelicnekopunch/go-sample/internal/usepackage"
	. "github.com/psychedelicnekopunch/go-sample/internal/usepackage"
	"github.com/psychedelicnekopunch/go-sample/internal/usepackage2"
)

func FuncSample() {
	fmt.Printf("func sample in main\n")
}

func main() {
	usepackage.Func()
	usepackageTest.Func()
	Func2()
	usepackage2.Func3()
	usepackage2.FuncSample()
	FuncSample()
}
