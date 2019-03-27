
package main


import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

func main() {
	fmt.Print(getExtension("The Velvet Underground & Nico.jpg"))// jpg

	fmt.Print(getExtension("The Velvet Underground & Nico.png"))// png

	fmt.Print(getExtension("The Velvet Underground & Nicojpg"))// ""
}


func getExtension(str string) string {

	// func MustCompile(str string) *Regexp
	// str = 正規表現
	// * バッククォート (`) で囲むのに注意。
	// * ダブルコーテーション (") でも囲めるけどいつも使う感じで使えないので注意
	// (?i) はよくある正規表現の "/xxxxxxxxxx/i" の i
	reg := regexp.MustCompile(`(?i)(\.jpg|\.jpeg|\.gif|\.png)$`)

	// func (re *Regexp) FindString(s string) string
	s := reg.FindString(str)
	if s == "" {
		return ""
	}

	// s[n:x] で n 番目から x 番目未満までの文字列を抜き出すことができる。
	// x が文字数以上の値だと panic が起きるので注意
	//
	// utf8.RuneCountInString(s) は文字列の文字数を返す
	// len(s) だと byte 数が返るので注意
	//
	// s[n:x] は下記コードと同じ処理だと思われる
	//
	// res := ""
	// for i := n; i < x; i++ {
	//     res = res + s[i]
	// }
	// return res
	return s[1:utf8.RuneCountInString(s)]
}