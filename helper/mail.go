package helper

import (
	"fmt"

	"gopkg.in/mail.v2"
)

// EmailConfig holds SMTP configuration
type EmailConfig struct {
	SMTPHost           string
	SMTPPort           int
	SenderMail         string
	SenderMailPassword string
	From               string
}

// EmailTemplateType defines the type of email
type EmailTemplateType string

const (
	OTPEmail          EmailTemplateType = "otp"
	NotificationEmail EmailTemplateType = "notification"
)

// SendEmail sends an email using the given configuration and template
func SendEmail(cfg EmailConfig, to string, subject string, templateType EmailTemplateType, data map[string]string) error {
	m := mail.NewMessage()
	m.SetHeader("From", cfg.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	// Choose template
	body := ""
	switch templateType {
	case OTPEmail:
		otp := data["otp"]
		body = fmt.Sprintf(`
        <html>
        <body>
            <div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #ddd; border-radius: 10px;">
                <h2 style="color: #4f46e5;">Your OTP Code</h2>
                <p>Use the following OTP to complete your verification:</p>
                <h1 style="color: #ef4444; font-size: 2rem;">%s</h1>
                <p>If you didn't request this, please ignore this email.</p>
            </div>
        </body>
        </html>
        `, otp)

	case NotificationEmail:
		message := data["message"]
		body = fmt.Sprintf(`
        <html>
        <body>
            <div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #ddd; border-radius: 10px;">
                <h2 style="color: #4f46e5;">Notification</h2>
                <p>%s</p>
            </div>
        </body>
        </html>
        `, message)

	default:
		body = "Hello, this is a default message."
	}

	m.SetBody("text/html", body)

	// SMTP Dialer
	d := mail.NewDialer(cfg.SMTPHost, cfg.SMTPPort, cfg.SenderMail, cfg.SenderMailPassword)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
