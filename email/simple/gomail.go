package main

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	m := gomail.NewMessage()

	// set e-mail sender
	m.SetHeader("From", "Anggit M Ginanjar <anggit.ginanjar.dev@gmail.com>")

	// set e-mail receiver
	m.SetHeader("To", "anggit@isi.co.id")

	// set e-mail subject
	m.SetHeader("Subject", "SIM InfinysCloud Web Client - User Invitation Link")

	// set e-mail body
	m.SetBody("text/plain", "Congratulations for your new account!")

	emailDialer := gomail.NewDialer("smtp.gmail.com", 587, "anggit.ginanjar.dev@gmail.com", "*1Systemadmin1*")
	emailDialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// send e-mail
	if err := emailDialer.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
