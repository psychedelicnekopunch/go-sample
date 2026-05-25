
package main


import (
	"fmt"
	"net/http"
)


func main() {
	fmt.Printf("%d %s\n", http.StatusConflict, http.StatusText(http.StatusConflict))
	fmt.Printf("%d %s\n", http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
}
