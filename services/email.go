package services

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// EmailObject defines email payload data
type EmailObject struct {
	To      string
	Body    string
	Subject string
}

var emailPass = []byte(os.Getenv("MAIL_SECRET"))

// SendMail method to send email to user
func SendMail(subject string, body string, to string, html string, name string) bool {
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))

	from := mail.NewEmail("Just Open it", os.Getenv("SENDGRID_FROM_MAIL"))
	_to := mail.NewEmail(name, to)
	plainTextContent := body
	htmlContent := html
	message := mail.NewSingleEmail(from, subject, _to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		return false
	} else {
		return true
	}
}
