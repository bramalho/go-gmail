package main

import (
	"log"
	"net/smtp"
)

func main() {
	from := "YOUR_EMAIL@gmail.com"
	passoword := "YOUR_PASSWORD"

	to := "to_email@mail.com"
	message := "It works!"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		smtp.PlainAuth("", from, passoword, "smtp.gmail.com"),
		from,
		[]string{to},
		[]byte(message),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Message Sent!")
}
