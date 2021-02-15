package main

import (
	"fmt"
	// "log"
	"os"
	"net/smtp"// https://golang.org/pkg/net/smtp/
)


func main() {

	from := "username@gmail.com"
	to := "test@gmail.com"

	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, "password", "smtp.gmail.com")

	err := send(auth, from, to, "テスト1")
	if err != nil {
		// fmt.Printf("%v\n", err)
		// log.Fatal(err)
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		return
	}

	err = send(auth, from, to, "テスト2")
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		return
	}

	fmt.Print("success")
}


func send(auth smtp.Auth, from string, to string, body string) (err error) {
	msg := []byte("" +
		"From: 送信した人 <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 件名 subject です\r\n\r\n" +
		body + "\r\n" +
	"")
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	return smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, msg)
}
