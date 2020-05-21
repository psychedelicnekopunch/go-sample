
package main


import (
	"fmt"
	"net/url"
)

func main() {
	s := "#apexチーター"
	fmt.Print(url.QueryEscape(s))
}
