package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	file, err := os.Open("images/sample.png")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer file.Close()


	copyTo, err := os.Create("images/sample_copy.png")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer copyTo.Close()

	_, err = io.Copy(copyTo, file)

	if err != nil {
		fmt.Printf(err.Error())
		if err := os.Remove("images/sample_copy.png"); err != nil {
			fmt.Printf(err.Error())
		}
		return
	}

	fmt.Printf("copied\n")
}
