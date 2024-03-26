package service

import (
	"awesomeProject3/api"
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

// SendEmail 添加SendEmail swagger注释
// @Summary		send email
// @Description	send email
// @Tags			email
// @Accept			json
// @Produce		json
// @Param    	targetEmail	query	string	true	"target email"
// @Param    	title	query	string	true	"title"
// @Param    	content	query	string	true	"content"
// @Success		200	{object}	api.User
// @Failure		400	{object}	error
// @Failure		404	{object}	error
// @Failure		500	{object}	error
// @Router			/email/send [get]
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
		authorEmail:    api.EmailAuthorEmail,
		authorPassword: api.EmailPassword,
	}
}
