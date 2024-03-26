package service

import (
	"awesomeProject3/api"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
	"strconv"
)

var location = "service/chat.go"

type Chat = api.Chat

type ChatService interface {
	// StartAChat / 增删改查
	StartAChat(title string, userId string, systemMessage string) error

	DeleteChat(id string) error

	DeleteAllChat(userId string) error

	GetChatDetail(id string) (Chat, error)

	GetChatList(chats *[]Chat, userId string) error

	SendMessage(chatId string, message string) error
}

type chatService struct{}

// StartAChat ShowAccount godoc
//	@Summary		start a chat
//	@Description	start a chat
//	@Tags			chat
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	api.User
//	@Failure		400	{object}	error
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Router			/chat/start [get]
func (c chatService) StartAChat(title string, id string, systemMessage string) error {
	log.Info().Msg("用户" + id + "开始一个" + title + "的聊天" + "location is :service/chat.go  StartAChat")
	if len(title) == 0 || title == "" {
		return errors.New("title can not be empty")
	}
	if len(systemMessage) == 0 {
		systemMessage = api.DefaultSystemMessage
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
	//打印
	log.Info().Msg("用户" + userId + "删除所有聊天" + "location is :service/chat.go  DeleteAllChat")
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
	log.Info().Msg("chat is :" + chat.Title + "location is :service/chat.go  GetChatDetail")
	return chat, nil
}

// GetChatList 获取聊天列表
func (c chatService) GetChatList(chats *[]Chat, userId string) error {
	Idint, _ := strconv.Atoi(userId)
	err := api.Db.Model(&Chat{}).Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return api.Db.Order("id desc")
	}).Where("user_id = ?", Idint).Find(&chats)
	if err.RowsAffected == 0 {
		return errors.New("chat not found")
	}
	return nil
}

// SendMessage 发送消息
func (c chatService) SendMessage(chatId string, message string) error {
	log.Info().Msg("chatId is :" + chatId + "location is :service/chat.go  SendMessage")

	chat := Chat{}
	err := api.Db.Model(&Chat{}).Preload("Messages").Find(&chat, chatId)

	if err.Error != nil || err.RowsAffected == 0 {
		return errors.New("chat not found")
	}

	//把chatId变为int
	cId, _ := strconv.Atoi(chatId)

	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleUser, Content: message})

	response := streamMessages(buildOpenAIMessages(&chat.Messages), int(chat.UserID), cId)

	//这里是在数据库更新用户发送的消息和AI回复的消息
	chat.Messages = append(chat.Messages, api.Message{Role: openai.ChatMessageRoleAssistant, Content: response})
	err = api.Db.Save(&chat)

	//如果失败
	if err.Error != nil {
		return errors.New("send message failed")
	}
	return nil
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
