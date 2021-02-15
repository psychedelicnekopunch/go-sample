package main

import (
	"fmt"
)

func toS(str string) *string {
	return &str
}

func main() {
	var str *string

	if str != nil {
		if *str != "" {
			fmt.Errorf("---1\n")
		}
	}
	if str != nil && *str != "" {
		fmt.Errorf("---1\n")
	}

	str = toS("")
	if str != nil && *str != "" {
		fmt.Printf("---2\n")
	}

	str = toS("test")
	if str != nil && *str != "" {
		fmt.Printf("---3\n")
	}

}
