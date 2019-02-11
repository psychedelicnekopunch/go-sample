
package main


import (
	"fmt"
	"crypto/sha256"// https://golang.org/pkg/crypto/sha256/
)


func main() {
	/**
	 * "crypto/sha256"
	 *
	 * func Sum256(data []byte) [Size]byte
	 * const Size = 32
	 */
	var data [sha256.Size]byte
	data = sha256.Sum256([]byte("test"))
	fmt.Print(data)
	fmt.Printf("\n")
	fmt.Printf("%x", data)

	var str string
	str = fmt.Sprintf("%x", data)

	fmt.Printf("\n")
	fmt.Printf(str)
}