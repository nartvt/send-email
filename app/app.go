package main

import (
	"fmt"
	"gopkg.in/mail.v2"
)

func main() {
	app := mail.NewMessage()

	app.SetHeader("From", "from@gmail.com")

	app.SetHeader("To", "toemail@gmail.com")

	app.SetHeader("Subject", "email send from go")

	app.SetBody("text/plain", "Here is email send from golang")

	email := mail.NewDialer("smtp.gmail.com", 587, "toemail@gmail.com", "password")
	//email.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := email.DialAndSend(app)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
