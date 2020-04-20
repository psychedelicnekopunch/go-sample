package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"
)

func main() {
	http := infrastructure.NewHttp()

	params := map[string]string{}
	params["sort"] = "created"
	params["direction"] = "desc"
	body, err := http.Get("https://api.github.com/users/psychedelicnekopunch/repos", params)

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	var res []struct {
		ID int `json:"id"`
		Name string `json:"name"`
		CreatedAt string `json:"created_at"`
	}

	dec := json.NewDecoder(strings.NewReader(body))
	dec.Decode(&res)
	// fmt.Print(res)

	for _, list := range res {
		fmt.Print(list)
		fmt.Print("\n")
	}
}
