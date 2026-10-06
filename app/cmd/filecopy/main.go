package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	file, err := os.Open("assets/img/sample.png")
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}
	defer file.Close()


	copyTo, err := os.Create("tmp/sample_copy.png")
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return
	}
	defer copyTo.Close()

	_, err = io.Copy(copyTo, file)

	if err != nil {
		fmt.Print(err.Error(), "\n")
		if err := os.Remove("tmp/sample_copy.png"); err != nil {
			fmt.Print(err.Error(), "\n")
		}
		return
	}

	fmt.Print("copied\n")
}
