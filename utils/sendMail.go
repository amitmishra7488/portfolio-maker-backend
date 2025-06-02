// utils/email.go
package utils

import (
	"crypto/tls"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(email, name, message, subject string) error {
	var emailTemplate string

	// Choose the email template based on the subject
	switch subject {
	case "Account Registration":
		emailTemplate = OtpVerificationTemplate(name, message)
	case "Reset Password":
		emailTemplate = OtpVerificationTemplate(name, message)
	case "Account Registration SuccessFully":
		emailTemplate = AccountRegistrationTemplate(name)
	default:
		return fmt.Errorf("invalid subject: %s", subject)
	}

	// Read SMTP configuration from environment variables
	from := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")
	smtpHost := os.Getenv("EMAIL_HOST")
	smtpPort := os.Getenv("EMAIL_PORT")

	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", fmt.Sprintf("Otp Verification for portfolio-builder - %s", subject))
	m.SetBody("text/html", emailTemplate)

	// Set up the SMTP server
	d := gomail.NewDialer(smtpHost, port, from, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Not recommended for production

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("✅ Email sent to:", email)
	return nil
}
