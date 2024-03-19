package service

import (
	"awesomeProject3/Constant"
	"awesomeProject3/api"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strconv"
)

type Chat = api.Chat

type ChatService interface {
	// StartAChat / 增删改查
	StartAChat(title string, userId string, systemMessage string) error

	DeleteChat(id string) error

	DeleteAllChat(userId string) error

	GetChatDetail(id string) (Chat, error)

	GetChatList(chats *[]Chat, userId string) error

	SendMessage(chatId string, message string) (string, error)
}

type chatService struct{}

// StartAChat 开始一个聊天
func (c chatService) StartAChat(title string, id string, systemMessage string) error {
	if len(title) == 0 || title == "" {
		return errors.New("title can not be empty")
	}
	if len(systemMessage) == 0 {
		systemMessage = Constant.DefaultSystemMessage
	}
	Idint, _ := strconv.Atoi(id)

	messages := []api.Message{
		{Role: openai.ChatMessageRoleSystem, Content: systemMessage},
	}

	chat := Chat{
		Title:         title,
		UserID:        uint(Idint),
		Messages:      messages,
		SystemMessage: systemMessage,
	}

	api.Db.Create(&chat)
	return nil
}

// DeleteChat delete chat
func (c chatService) DeleteChat(id string) error {
	chat := Chat{}
	api.Db.Find(&chat, id)
	if chat.ID == 0 {
		return errors.New("chat not found")
	}
	api.Db.Delete(&Chat{}, id)
	return nil
}

// DeleteAllChat delete all chat
func (c chatService) DeleteAllChat(userId string) error {
	Idint, _ := strconv.Atoi(userId)
	err := api.Db.Where("user_id = ?", Idint).Delete(&Chat{})
	if err.Error != nil {
		return errors.New("delete chat failed")
	}
	return nil
}

// GetChatDetail 获取聊天详情
func (c chatService) GetChatDetail(id string) (Chat, error) {
	chat := Chat{}
	err := api.Db.Model(&Chat{}).Preload("Messages").Find(&chat, id)
	if chat.ID == 0 || err.RowsAffected == 0 {
		return chat, errors.New("chat not found")
	}
	fmt.Println(chat, "location:service/chat.go GetChatDetail")
	return chat, nil
}

// GetChatList 获取聊天列表
func (c chatService) GetChatList(chats *[]Chat, userId string) error {
	Idint, _ := strconv.Atoi(userId)
	err := api.Db.Model(&Chat{}).Preload("Messages").Where("user_id = ?", Idint).Find(&chats)
	if err.RowsAffected == 0 {
		return errors.New("chat not found")
	}
	return nil
}

// SendMessage 发送消息
func (c chatService) SendMessage(chatId string, message string) (string, error) {
	fmt.Println("chatId is :"+chatId, "message is :"+message, "    location is :service/chat.go  SendMessage")

	chat := Chat{}
	err := api.Db.Model(&Chat{}).Preload("Messages").Find(&chat, chatId)

	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleUser, Content: message})

	if err.Error != nil || err.RowsAffected == 0 {
		return "", errors.New("chat not found")
	}

	response := sendMessage(buildOpenAIMessages(&chat.Messages))

	//这里是在数据库更新用户发送的消息和AI回复的消息
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleAssistant, Content: response})
	err = api.Db.Save(&chat)

	//如果失败
	if err.Error != nil {
		return "", errors.New("send message failed")
	}
	return response, nil
}

// SaveAIResponse 保存ai回复 到数据库
func (c chatService) SaveAIResponse(chatId string, content string) {
	chat := Chat{}
	api.Db.Find(&chat, chatId)
	if chat.ID == 0 {
		return
	}
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleAssistant, Content: content})
	err := api.Db.Save(&chat)
	if err.Error != nil {
		return
	}
	return
}

func NewChatService() ChatService {
	return &chatService{}
}
