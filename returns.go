package main

import "fmt"

type BadRequestError struct {
	Message string
	Code int
}


func (req BadRequestError) Error() string {
	return fmt.Sprintf("%v: %v", req.Message, req.Code)
}


func Returns(num int) (int, error) {
	if num < 0 {
		return -1, BadRequestError{ "num is required", 400 }
	}
	return num, nil
}

func main() {
	res, err := Returns(1);
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Printf("%v", res)
}