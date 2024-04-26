package service

import (
	"awesomeProject3/config"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type EmailService interface {
	SendEmail(targetUserId string, title string, content string) error
}

type emailService struct {
	authorEmail    string
	authorPassword string
	title          string
	template       string
}

func (e emailService) SendEmail(targetEmail string, title string, content string) error {
	fmt.Println("send email to "+targetEmail, "   location is :service/email.go  SendEmail")
	newEmail := email.NewEmail()
	newEmail.From = "Zing  <" + e.authorEmail + ">"
	newEmail.To = []string{targetEmail}
	newEmail.Subject = title
	newEmail.Text = []byte(content)
	err := newEmail.Send("smtp.qq.com:25", smtp.PlainAuth("", e.authorEmail, e.authorPassword, "smtp.qq.com"))
	if err != nil {
		return err
	}
	return nil
}

func NewEmailService() EmailService {
	return &emailService{
		authorEmail:    config.GetConfig().Email.EmailAuthorEmail,
		authorPassword: config.GetConfig().Email.EmailPassword,
		title:          config.GetConfig().Email.EmailTitle,
		template:       config.GetConfig().Email.EmailTemplate,
	}
}
