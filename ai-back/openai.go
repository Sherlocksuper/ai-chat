package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
)

var token = "sk-aIYgLMM0SGNroc9n2270Ed56Af2f403bAb652b77C0F0BbA6"
var model = openai.GPT3Dot5Turbo
var maxTokens = 400
var useStream = true
var client = openai.NewClient(token)

var messages []openai.ChatCompletionMessage

func init() {

}

// / Add system message
func addSystemMessage(content string) {
	message := buildSystemMessage(content)
	saveMessageToHistory(message)
}

// / 发送message
func sendMessage(message openai.ChatCompletionMessage) {
	saveMessageToHistory(message)
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     model,
		MaxTokens: maxTokens,
		Messages:  messages,
		Stream:    useStream,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	totalMessages := ""
	for {
		response, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			break
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		totalMessages += response.Choices[0].Delta.Content
		fmt.Print(response.Choices[0].Delta.Content)
	}
	saveAIResponse(totalMessages)
}

// / 用户发送消息
func userSendMessage(content string) {
	message := buildUserMessage(content)
	saveMessageToHistory(message)
	sendMessage(message)
}

// / 保存ai的回复
func saveAIResponse(content string) {
	message := buildAIResponse(content)
	saveMessageToHistory(message)
}

// / 保存消息到历史记录
func saveMessageToHistory(message openai.ChatCompletionMessage) {
	messages = append(messages, message)
}

// / build message
func buildUserMessage(content string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	}
}

// / build system message
func buildSystemMessage(content string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: content,
	}
}

// / build ai response
func buildAIResponse(content string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	}
}
