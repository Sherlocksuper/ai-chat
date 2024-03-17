package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:root@tcp(127.0.0.1:3306)/ai?charset=utf8mb4&parseTime=True&loc=Local"
var Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

var Token = "sk-aIYgLMM0SGNroc9n2270Ed56Af2f403bAb652b77C0F0BbA6"
var Model = openai.GPT3Dot5Turbo
var MaxTokens = 400
var UseStream = true
var Client = openai.NewClient(Token)

func init() {

}

// 一些返回的code
const (
	SUCCESS     = 200
	FAIL        = 400
	NOTFOUND    = 404
	SERVERERROR = 500
)

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
