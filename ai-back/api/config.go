package api

import (
	"github.com/sashabaranov/go-openai"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db, _ = gorm.Open(mysql.Open(DSN), &gorm.Config{})

// openai go服务
const (
	DSN         = "root:root@tcp(localhost:3306)/ai?charset=utf8mb4&parseTime=True&loc=Local"
	OPENAITOKEN = "sk-aIYgLMM0SGNroc9n2270Ed56Af2f403bAb652b77C0F0BbA6"
	MODEL       = openai.GPT3Dot5Turbo
	MAXTOKENS   = 2000
)

// email go服务 与 redis go 服务放到各自的service里面
