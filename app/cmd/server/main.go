package main


import (
	"fmt"
	"net/http"
)


// $ go run ./server.go
// http://localhost:8080/ にアクセス
func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Print(err.Error())
	}
}


func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello, world")
}
