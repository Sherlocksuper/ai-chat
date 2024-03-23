package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/service"
	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	emailService service.EmailService
}

func NewEmailHandler(emailService service.EmailService) *EmailHandler {
	return &EmailHandler{
		emailService: emailService,
	}
}

// SendEmail 发送邮件 POST
func (f *EmailHandler) SendEmail(c *gin.Context) {
	var email api.Email
	err := c.BindJSON(&email)

	if err != nil {
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
	}

	err = f.emailService.SendEmail(email.TargetEmail, email.Title, email.Content)

	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "发送成功", nil))
}
