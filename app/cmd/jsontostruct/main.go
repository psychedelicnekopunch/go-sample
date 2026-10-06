package main

import (
	"fmt"
	"encoding/json"
	"strings"
)


func main() {

	jsonData := `{"result":"success","hello":"world"}`

	var persedJson struct {
		Result string `json:"result"`
		Test int `json:"test"`
		Hello string `json:"hello"`
	}

	dec := json.NewDecoder(strings.NewReader(jsonData))
	dec.Decode(&persedJson)

	fmt.Print(persedJson, "\n")
}
