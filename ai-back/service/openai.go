package service

import (
	"awesomeProject3/api"
	"awesomeProject3/ws"
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
)

var client = openai.NewClient(api.Token)

func init() {

}

// /流式接收信息
func streamMessages(messages []openai.ChatCompletionMessage, userId int, chatId int) string {
	c := openai.NewClient(api.Token)
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     api.Model,
		MaxTokens: api.MaxTokens,
		Messages:  messages,
		Stream:    true,
	}

	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Println(err.Error() + `,location:openai.go  streamMessages`)
		return ""
	}

	defer stream.Close()

	totalResponse := ""

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			break
		}

		if err != nil {
			fmt.Println("\nStream error: %v\n", err)
			break
		}

		totalResponse += response.Choices[0].Delta.Content
		fmt.Println(response.Choices[0].Delta.Content)

		ws.SendMsg(userId, ws.WsReMessage{
			Type: ws.CHAT_MESSAGE,
			Content: ws.ChatResContent{
				UserId:  userId,
				ChatId:  chatId,
				Message: response.Choices[0].Delta.Content,
			},
		})
	}

	return totalResponse
}

func buildOpenAIMessages(messages *[]api.Message) []openai.ChatCompletionMessage {
	var openAIMessages []openai.ChatCompletionMessage

	println("messages is : location:service/openai.go  buildOpenAIMessages")
	for _, message := range *messages {
		println(message.Content)
		openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	return openAIMessages
}
