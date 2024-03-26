package service

import (
	"awesomeProject3/api"
	"awesomeProject3/ws"
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"io"
)

func init() {

}

// /流式接收信息
func streamMessages(messages []openai.ChatCompletionMessage, userId int, chatId int) string {
	config := openai.DefaultConfig(api.OPENAITOKEN)
	config.BaseURL = api.BASEURL
	c := openai.NewClientWithConfig(config)

	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     api.MODEL,
		MaxTokens: api.MAXTOKENS,
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
			log.Info().Msg("Stream finished")
			break
		}

		if err != nil {
			log.Error().Msg("Stream error: " + err.Error())
			break
		}

		totalResponse += response.Choices[0].Delta.Content

		ws.SendMsg(userId, ws.WsReMessage{
			Type: ws.CHAT_MESSAGE,
			Content: ws.ChatResContent{
				UserId:  userId,
				ChatId:  chatId,
				Message: response.Choices[0].Delta.Content,
			},
		})
	}

	//打印 ai回复chatId + 消息内容
	log.Info().Msg("response to user :" + string(userId) + "  message is :" + totalResponse + "location is :service/openai.go  streamMessages")

	return totalResponse
}

func buildOpenAIMessages(messages *[]api.Message) []openai.ChatCompletionMessage {
	var openAIMessages []openai.ChatCompletionMessage

	log.Info().Msg("messages is :" + fmt.Sprint(messages) + "location is :service/openai.go  buildOpenAIMessages")
	for _, message := range *messages {
		openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	return openAIMessages
}
