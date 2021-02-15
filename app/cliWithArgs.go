
package main


import (
	"flag"
	"fmt"
	// "os"
)


func main() {
	flag.Parse()
	args := flag.Args()
	// fmt.Print(args)
	// fmt.Print(os.Args)

	for _, v := range args {
		fmt.Printf("%s\n", v)
	}
}
