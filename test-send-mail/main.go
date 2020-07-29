package main

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/mail.v2"
	//"github.com/go-mail/mail"
)

func main() {

	fmt.Println("Try sending mail...")

	// local smtp server
	d := mail.NewDialer("localhost", 2525, "auth@mail.com", "S3CR3T")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// gmail
	//d := mail.NewDialer("smtp.gmail.com", 587, "username@gmail.com", "PASSW0RD")

	m := mail.NewMessage()

	m.SetHeader("From", "john@mail.com")
	m.SetHeader("To", "alice@gmail.com")
	m.SetHeader("Subject", "Test ทดสอบ テスト !!!")
	m.SetBody("text/plain", "This is a test ทดสอบ テスト")

	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Failed sending mail")
		panic(err)
	}

	fmt.Println("Mail sent without error")
}
