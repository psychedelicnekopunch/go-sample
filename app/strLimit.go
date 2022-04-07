
package main


import (
	"fmt"
	// "strconv"
	"unicode/utf8"
)


func limit(str string) string {
	return string([]rune(str)[:30])
	// return string(utf8.RuneCountInString(str)[:15])
}


func main() {

	str := "abc, 日本語です。"
	fmt.Printf("%s (%d Words)(%d Byte)\n", str, utf8.RuneCountInString(str), len(str))
	str = limit(str)
	fmt.Printf("%s (%d Words)(%d Byte)\n", str, utf8.RuneCountInString(str), len(str))


	str2 := "abc, 日本語です。 文字数制限LIMIT abc, 日本語です。 文字数制限LIMIT abc, 日本語です。 文字数制限LIMIT"
	fmt.Printf("%s (%d Words)(%d Byte)\n", str2, utf8.RuneCountInString(str2), len(str2))
	str2 = limit(str2)
	fmt.Printf("%s (%d Words)(%d Byte)\n", str2, utf8.RuneCountInString(str2), len(str2))
}
