package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const (
	CONFIG_SMTP_HOST     = "smtp.gmail.com"
	CONFIG_SMTP_PORT     = 587
	CONFIG_SENDER_NAME   = "Anggit M Ginanjar <anggit.ginanjar.dev@gmail.com>"
	CONFIG_AUTH_EMAIL    = "anggit.ginanjar.dev@gmail.com"
	CONFIG_AUTH_PASSWORD = "*1Systemadmin1*"
)

func main() {
	to := []string{"anggit@isi.co.id"}
	cc := []string{}
	subject := "InfinysCloud web client user invitation"
	message := "Congratulations for your new account! Recovery link: https://infinyscloud.com/recovery_link"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message
	//body := fmt.Sprintf("From: %s \n To: %s \n Cc: %s \n Subject: %s \n %s", CONFIG_SENDER_NAME, strings.Join(to, ","), strings.Join(cc, ","), subject, message)

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
