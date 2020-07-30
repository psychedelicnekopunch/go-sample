package main

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron/v3"
)

func main() {

	c := cron.New()
	c.AddFunc("@every 2s", func() {
		call()
	})
	c.AddFunc("@every 3s", call2)
	c.Start()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf(err.Error())
	}
}


func call() {
	fmt.Printf("tset\n")
}


func call2() {
	fmt.Printf("tset 2\n")
}
