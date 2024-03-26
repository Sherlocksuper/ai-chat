package api

import (
	"github.com/sashabaranov/go-openai"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db, _ = gorm.Open(mysql.Open(DSN), &gorm.Config{})

// openai go服务
const (
	DSN = "root:root@tcp(localhost:3306)/ai?charset=utf8mb4&parseTime=True&loc=Local"
	//OPENAITOKEN          = "sk-aIYgLMM0SGNroc9n2270Ed56Af2f403bAb652b77C0F0BbA6"
	//BASEURL              = "https://hk.xty.app/v1"
	//MODEL       = openai.GPT3Dot5Turbo
	OPENAITOKEN = "sk-UR1Vea04XpkoPi1T071cB5A4FdF94dFf89Ba7933DaB42005"
	BASEURL     = "https://izepkqss.cloud.sealos.io"
	//BASEURL              = "https://api.openai.com/v1"
	MODEL                = openai.GPT4
	MAXTOKENS            = 2000
	DefaultSystemMessage = "你是我的ai助手，请帮助我解决问题"
)

// email go服务
const (
	EmailAuthorEmail = "1075773551@qq.com"
	EmailPassword    = "snghbtzvldxoidab"
	EmailTitle       = "注册验证码"
	EmailTemplate    = "您的注册验证码为：%s, 请在5分钟内完成注册"
)

// redis服务
const (
	RedisAddress  = "localhost:6379"
	RedisPassword = ""
	RedisDb       = 0
)
