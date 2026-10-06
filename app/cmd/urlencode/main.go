package main

import (
	"fmt"
	"net/url"
)


func main() {
	str := "?q=t%e%s%t&❤️ 絵+文¥字_ やゔ/ぁぴぼ _ 空白 \" ' ? ! .xlsx .png ♡"
	fmt.Print(str, "\n\n")
	fmt.Print(url.PathEscape(str), "\n\n")
	fmt.Print(url.QueryEscape(str), "\n\n")
}
