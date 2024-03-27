package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
)

type ChatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(chatService service.ChatService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
	}
}

// StartAChat 开始一个聊天 POST
func (f *ChatHandler) StartAChat(c *gin.Context) {
	var chat service.Chat
	err := c.BindJSON(&chat)

	if err != nil {
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
	}

	err = f.chatService.StartAChat(chat.Title, strconv.Itoa(int(chat.UserID)), chat.SystemMessage)

	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "创建成功", nil))
}

// DeleteChat 删除一个聊天 GET  chatId
func (f *ChatHandler) DeleteChat(c *gin.Context) {
	id := c.Query("chatId")
	err := f.chatService.DeleteChat(id)

	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "删除成功", nil))
}

// 删除所有Chat
func (f *ChatHandler) DeleteAllChat(c *gin.Context) {
	userId := c.Query("userId")
	println("userId is :"+userId, "location is :handler/chat.go")
	if userId == "" {
		c.JSON(200, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err := f.chatService.DeleteAllChat(userId)

	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "删除成功", nil))
}

// GetChatDetail 获取聊天详情 GET
func (f *ChatHandler) GetChatDetail(c *gin.Context) {
	id := c.Query("id")
	chat, err := f.chatService.GetChatDetail(id)

	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "获取成功", chat))
}

// GetChatList 获取聊天列表 GET
func (f *ChatHandler) GetChatList(c *gin.Context) {
	userId := c.Query("userId")

	var chats []service.Chat
	f.chatService.GetChatList(&chats, userId)

	c.JSON(200, api.M(api.SUCCESS, "获取成功", chats))
}

// SendMessage 发送消息 POST
func (f *ChatHandler) SendMessage(c *gin.Context) {
	var message api.Message
	err := c.BindJSON(&message)

	log.Info().Msg("message is :" + fmt.Sprint(message.Content) + "location is :handler/chat.go  SendMessage")
	if err != nil {
		c.JSON(200, api.M(api.FAIL, "参数错误", nil))
		return
	}

	//发送消息
	err = f.chatService.SendMessage(c, strconv.Itoa(message.ChatID), message.Content)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
}
