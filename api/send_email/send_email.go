package send_email

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendEmail sends an email to the specified recipient using the SMTP relay
func SendEmail(to, subject, body string) error {
	from := os.Getenv("EMAIL_ADDRESS")
	smtpHost := "smtp" // This should match the service name in docker-compose.yml
	smtpPort := "25"

	// Message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, nil, from, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}
