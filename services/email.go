package services

import (
	"fmt"
	"log"
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
func SendMail() {
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))

	from := mail.NewEmail("Test User", os.Getenv("SENDGRID_FROM_MAIL"))
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Test User", "ewachira254@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
