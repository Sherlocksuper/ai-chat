package service

import (
	"awesomeProject3/api"
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var token = "sk-aIYgLMM0SGNroc9n2270Ed56Af2f403bAb652b77C0F0BbA6"
var model = openai.GPT3Dot5Turbo
var maxTokens = 1000
var useStream = true
var client = openai.NewClient(token)

func init() {

}

// / 发送message
func sendMessage(messages []openai.ChatCompletionMessage) string {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     model,
		MaxTokens: maxTokens,
		Messages:  messages,
		Stream:    false,
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		req,
	)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	//循环打印resp.Choices
	fmt.Println("resp.Choices is below ,location:openai.go  sendMessage")
	for _, message := range resp.Choices {
		fmt.Println("role is :" + message.Message.Role + "  message is :" + message.Message.Content)
	}

	//返回最后一条消息
	return resp.Choices[len(resp.Choices)-1].Message.Content
}

func buildOpenAIMessages(messages *[]api.Message) []openai.ChatCompletionMessage {
	var openAIMessages []openai.ChatCompletionMessage

	for _, message := range *messages {
		openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	return openAIMessages
}
