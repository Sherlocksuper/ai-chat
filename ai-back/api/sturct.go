package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Email    string `json:"email"`
	Chats    []Chat `json:"chats"`
}

type Chat struct {
	gorm.Model
	Title         string    `json:"title"`
	UserID        uint      `json:"userId"`
	SystemMessage string    `json:"systemMessage"`
	Messages      []Message `json:"messages"`
}

type Message struct {
	gorm.Model
	ChatID  int    `json:"chatId"`
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Version struct {
	gorm.Model
	Version      string `json:"version" gorm:"unique"`
	Introduction string `json:"introduction"`
	Enable       bool   `json:"enable"`
	DownloadUrl  string `json:"downloadUrl"`
}

type Email struct {
	gorm.Model
	TargetEmail string `json:"targetEmail"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}

type Prompt struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Function string `json:"function"`
}

// 一些返回的code
type ReturnMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Func struct {
	Url    string `json:"url"`
	Method string `json:"method"`
	Action func(c *gin.Context)
}
