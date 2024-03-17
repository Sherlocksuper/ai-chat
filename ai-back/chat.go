package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

type mes = openai.ChatCompletionMessage

func sendMessageToAI(c *gin.Context) {

	//接收json数据 message
	var message mes
	err := c.BindJSON(&message)
	fmt.Println(message)
	if err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	c.JSON(200, gin.H{"message": "success"})
}
