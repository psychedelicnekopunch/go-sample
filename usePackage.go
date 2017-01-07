package main

import (
	"fmt"
	"go-sample/models"
	modelsSample "go-sample/models"
	. "go-sample/models"
	"go-sample/models2"
)

func FuncSample() {
	fmt.Printf("func sample in main\n")
}

func main() {
	models.Func()
	modelsSample.Func()
	Func2()
	models2.Func3()
	models2.FuncSample()
	FuncSample()
}

