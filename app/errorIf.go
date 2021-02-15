package main

import (
	"fmt"
	"errors"
)

func success() (mes string, err error) {
	return "success", nil
}

func failed() (mes string, err error) {
	return "failed", errors.New("error message")
}


func ptn1() {
	fmt.Printf("======= ptn1 ========\n")

	var mes string
	var mes2 string
	var err error

	if mes, err = success(); err != nil {
		fmt.Printf(err.Error() + "\n")
	}
	fmt.Printf(mes + "\n")


	if mes2, err = failed(); err != nil {
		fmt.Printf(err.Error() + "\n")
	}
	fmt.Printf(mes2 + "\n\n")
}


func ptn2() {
	fmt.Printf("======= ptn2 ========\n")

	mes, err := success()

	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}
	fmt.Printf(mes + "\n")


	mes2, err := failed()

	if err != nil {
		fmt.Printf(err.Error() + "\n")
	}
	fmt.Printf(mes2 + "\n\n")
}


func main() {
	ptn1()
	ptn2()
}
