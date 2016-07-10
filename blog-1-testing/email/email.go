package email

import "gopkg.in/gomail.v2"

var (
	NewSignupSubject = "New user signed up"
	InviteSubject    = "You are being invited to awesome product"
)

type EmailSender interface {
	SendEmail(to []string, from, subject string, body []byte) error
}

// Amazon SES email sender
type SESEmailSender struct {
	smtpUsername string
	smtpPassword string
	host         string
	port         int
}

func (s *SESEmailSender) SendEmail(to []string, from, subject string, body []byte) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", string(body))
	d := gomail.NewPlainDialer(s.host, s.port, s.smtpUsername, s.smtpPassword)
	return d.DialAndSend(m)
}

var Sender EmailSender
var fromEmail = "product@thecompany.com"

func init() {
	Sender = &SESEmailSender{
		smtpUsername: "748do0fl23xfe89",
		smtpPassword: "xxxxxxxx",
		host:         "email-smtp.us-west-2.amazonaws.com",
		port:         587,
	}
}

func SendNewSignupEmailToAdmin(adminEmail string, body []byte) error {
	if err := Sender.SendEmail([]string{adminEmail}, fromEmail, NewSignupSubject, body); err != nil {
		return err
	}
	return nil
}

func SendInvitationEmail(to []string, body []byte) error {
	if err := Sender.SendEmail(to, fromEmail, InviteSubject, body); err != nil {
		return err
	}
	return nil
}
