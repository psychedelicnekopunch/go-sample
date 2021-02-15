package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"
)


type Mail struct {
	Auth smtp.Auth
	Addr string
	From mail.Address
}


func NewMail() *Mail {
	c := infrastructure.NewConfig()
	return &Mail{
		Auth: smtp.PlainAuth("", c.Mail.Gmail.Address, c.Mail.Gmail.Password, c.Mail.Gmail.Host),
		Addr: c.Mail.Gmail.Addr,
		From: mail.Address{ Name: c.Mail.Gmail.Name, Address: c.Mail.Gmail.Address },
	}
}


func (m *Mail) writeString(b *bytes.Buffer, s string) *bytes.Buffer {
	_, err := b.WriteString(s)
	if err != nil {
		fmt.Print(err.Error())
	}
	return b
}


/**
 * go で utf8 メールを送信
 * https://qiita.com/yamasaki-masahide/items/a9f8b43eeeaddbfb6b44
 *
 * サンプルコードほぼ丸パクリ
 */

// サブジェクトを MIME エンコードする
func (m *Mail) encodeSubject(subject string) string {
	// UTF8 文字列を指定文字数で分割する
	b := bytes.NewBuffer([]byte(""))
	strs := []string{}
	length := 13
	for k, c := range strings.Split(subject, "") {
		b.WriteString(c)
		if k%length == length-1 {
			strs = append(strs, b.String())
			b.Reset()
		}
	}
	if b.Len() > 0 {
		strs = append(strs, b.String())
	}
	// MIME エンコードする
	b2 := bytes.NewBuffer([]byte(""))
	b2.WriteString("Subject:")
	for _, line := range strs {
		b2.WriteString(" =?utf-8?B?")
		b2.WriteString(base64.StdEncoding.EncodeToString([]byte(line)))
		b2.WriteString("?=\r\n")
	}
	return b2.String()
}


// 本文を 76 バイト毎に CRLF を挿入して返す
func (m *Mail) encodeBody(body string) string {
	b := bytes.NewBufferString(body)
	s := base64.StdEncoding.EncodeToString(b.Bytes())
	b2 := bytes.NewBuffer([]byte(""))
	for k, c := range strings.Split(s, "") {
		b2.WriteString(c)
		if k % 76 == 75 {
			b2.WriteString("\r\n")
		}
	}
	return b2.String()

}


func (m *Mail) Send(to string, subject string, body string) (err error) {

	msg := bytes.NewBuffer([]byte(""))
	msg = m.writeString(msg, "From: " + m.From.String() + "\r\n")
	msg = m.writeString(msg, "To: " + to + "\r\n")
	msg = m.writeString(msg, "Bcc: " + m.From.String() + "\r\n")
	// msg = m.writeString(msg, "Subject: " + subject + "\r\n")
	msg = m.writeString(msg, m.encodeSubject(subject))
	msg = m.writeString(msg, "MIME-Version: 1.0\r\n")
	msg = m.writeString(msg, "Content-Type: text/plain; charset=\"utf-8\"\r\n")
	msg = m.writeString(msg, "Content-Transfer-Encoding: base64\r\n")
	msg = m.writeString(msg, "\r\n")

	// msg = m.writeString(msg, base64.StdEncoding.EncodeToString([]byte(body)) + "\r\n")
	msg = m.writeString(msg, m.encodeBody(body))

	smtp.SendMail(m.Addr, m.Auth, m.From.Address, []string{to}, msg.Bytes())

	fmt.Print(msg, "\n")
	fmt.Print(body, "\n")

	return nil
}


func main() {

	to := "pipoist@gmail.com"
	subject := "件名Testです"
	body := `test
this is a body.
日本語も文字化けしないと思います。
テストです。`

	mail := NewMail()
	err := mail.Send(to, subject, body)

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Print("success")
}
