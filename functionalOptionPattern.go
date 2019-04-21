package main

import "fmt"


type Sample struct {
	Val string
}

func setVal(value string) string {
	if value == "" {
		return "temp value"
	}
	return value
}


func NewSample(value string) *Sample {
	s := new(Sample)
	s.Val = value
	return s
}

func main() {
	sample := NewSample(setVal(""))
    fmt.Print(sample)

    sample2 := NewSample(setVal("test"))
    fmt.Print(sample2)
}


