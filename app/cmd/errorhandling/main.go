package main

import (
	"fmt"
	"errors"
)

type BadRequestError struct {
	Message string
	Code int
}

type NotFoundError struct {
	Message string
	Code int
}

type ServerError struct {
	Message string
	Code int
}

func (req BadRequestError) Error() string {
	return fmt.Sprintf("%v: %v", req.Message, req.Code)
}

func (req NotFoundError) Error() string {
	return fmt.Sprintf("%v: %v", req.Message, req.Code)
}

func (req ServerError) Error() string {
	return fmt.Sprintf("%v: %v", req.Message, req.Code)
}


func StatusCodes(num int) (int, error) {
	switch num {
	case 200:
		return num, nil
	case 400:
		return num, BadRequestError{ "num is bad request", num }
	case 404:
		return num, NotFoundError{ "num is not found", num }
	case 500:
		return num, ServerError{ "num is server error", num }
	default:
		return num, errors.New("unknown error")
	}
}


func SendResponse(res int, err error) {
	if err != nil {
		// https://go.dev/tour/methods/16
		switch err.(type) {
		case BadRequestError:
			fmt.Printf("BadRequestError<%T>: %v\n", err, err)
		case NotFoundError:
			fmt.Printf("NotFoundError<%T>: %v\n", err, err)
		case ServerError:
			fmt.Printf("ServerError<%T>: %v\n", err, err)
		default:
			fmt.Printf("Error<%T>: %v\n", err, err)
		}
		return
	}
	if res == 200 {
		fmt.Printf("Success: %v\n", res)
	} else {
		fmt.Printf("Unknown: %v\n", res)
	}
}


func main() {
	res, err := StatusCodes(404);
	SendResponse(res, err)
}
