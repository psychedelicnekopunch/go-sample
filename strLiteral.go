
package main


import (
	"fmt"
)


func main() {

	param := "success"
	i := 10

	fmt.Print(fmt.Sprintf(
`test
%s, %d
`,
	param,
	i))
}
