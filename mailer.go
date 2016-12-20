package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

func SendMail() {
	ch := make(chan *gomail.Message)

	go func() {
		d := gomail.NewDialer("smtp.example.org", 587, "user", "123456")

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					return
				}

				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}

				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}

			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}

					open = false
				}
			}
		}
	}()

	// TEST
	for i := 0; i < 10; i++ {
		m := composeMessage(i)
		ch <- m
	}

	// Close the channel to stop the mail daemon.
	close(ch)
}

func composeMessage(i int) *gomail.Message {
	m := gomail.NewMessage()

	m.SetHeader("From", "example@example.org")
	m.SetHeader("To", "example@example.org")
	m.SetHeader("Subject", fmt.Sprintf("Login request #%d from Nogopass!", i))
	m.SetBody("text/html", "Hello <b>tester!</b><br />Please click on the following link to ")

	return m
}
