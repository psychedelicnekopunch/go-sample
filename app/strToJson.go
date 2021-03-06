package main


import (
	"encoding/json"
	"fmt"
	"strings"
)


func main() {

	str := `{"id":1,"title":"test","description":"概要です","note":"注意事項です","url":null}`

	var res struct {
		ID int `json:"id"`
		Title string `json:"title"`
		Description string `json:"description"`
		Note string `json:"note"`
		URL string `json:"url"`
	}

	dec := json.NewDecoder(strings.NewReader(str))
	dec.Decode(&res)

	fmt.Printf(fmt.Sprintf("\nJSON Decoded ====\n%d\n%s\n%s\n%s\n%s\n", res.ID, res.Title, res.Description, res.Note, res.URL))
	// JSON Decoded ====
	// 1
	// test
	// 概要です
	// 注意事項です

	var lists []struct {
		ID int `json:"id"`
		Title string `json:"title"`
	}

	strs := `[{"id":1,"title":"test"},{"id":2,"title":"test2"},{"id":3,"title":"test3"}]`
	dec2 := json.NewDecoder(strings.NewReader(strs))
	dec2.Decode(&lists)
	fmt.Print(lists)
}
