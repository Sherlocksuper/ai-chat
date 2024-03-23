package service

import (
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
}

func (e emailService) SendEmail(targetEmail string, title string, content string) error {
	fmt.Println("send email to "+targetEmail, "   location is :service/email.go  SendEmail")
	newEmail := email.NewEmail()
	newEmail.From = "dj <" + e.authorEmail + ">"
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
		authorEmail:    "1075773551@qq.com",
		authorPassword: "snghbtzvldxoidab",
	}
}