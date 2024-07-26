package jobs

import (
	"crypto/tls"
	"log"
	"strconv"

	"github.com/whatever/your/app/name/is/model"
	"gopkg.in/gomail.v2"
)

var from string = "tony starq"
var host string = "whateveryouremailhostis.com"
var port string = "411"
var username string = "email-username"
var password string = "email-password"

type EmailJobData struct {
	User    model.User
	Subject string
	Body    string
}

func SendEmail(data EmailJobData) (result string, message string, error error) {

	to := []string{data.User.Email}
	log.Println(to[0])

	m := gomail.NewMessage()
	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", data.User.Email)

	// Set E-Mail subject
	m.SetHeader("Subject", data.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", data.Body)

	intPort, _ := strconv.Atoi(port)
	// Settings for SMTP server
	d := gomail.NewDialer(host, intPort, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return "Failed", "Oops!! Something went wrong", err

	}

	log.Println("Email sent successfully!")
	return "Success", "Email Successfully Sent To " + data.User.Email, nil

}
