package service

import (
	"awesomeProject3/api"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
)

func init() {

}

// /流式接收信息
func streamMessages(messages []openai.ChatCompletionMessage, ginCon *gin.Context) string {
	//设置openai的配置
	config := openai.DefaultConfig(api.OPENAITOKEN)
	config.BaseURL = api.BASEURL
	c := openai.NewClientWithConfig(config)

	//设置网络writer以及
	w := ginCon.Writer
	w.WriteHeader(200)

	//开启stream
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     api.MODEL,
		MaxTokens: api.MAXTOKENS,
		Messages:  messages,
		Stream:    true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)

	if err != nil {
		log.Error().Msg("Stream error: " + err.Error())
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

		w.Write([]byte(response.Choices[0].Delta.Content))
		w.(http.Flusher).Flush()
		fmt.Print(response.Choices[0].Delta.Content)
	}

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
