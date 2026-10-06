
package main


import (
	"fmt"
	"net/url"
)

func main() {
	s := "#カレーOTP カツ丼食いてえ"
	fmt.Print(url.QueryEscape(s), "\n")
}
