package main

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron"
)

func main() {

	c := cron.New()
	c.AddFunc("@every 2s", func() {
		fmt.Printf("cron tset\n")
	})
	c.Start()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Printf(err.Error())
	}
}