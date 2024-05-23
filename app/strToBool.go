package main


import (
	"fmt"
	"strconv"
)


func main() {

	// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value returns an error.
	fmt.Printf("%T型: %v\n", parseBool("true"), parseBool("true"))
	fmt.Printf("%T型: %v\n", parseBool("false"), parseBool("false"))
	fmt.Printf("%T型: %v\n", parseBool("true"), parseBool("2"))
}


func parseBool(str string) bool {
	t, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Print("ERROR: ", err.Error(), "\n")
	}
	return t
}
