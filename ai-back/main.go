package main

import (
	"awesomeProject3/api"
	"awesomeProject3/handler"
	"awesomeProject3/service"
	"awesomeProject3/ws"
	"github.com/gin-gonic/gin"
)

type request struct {
	Type string `json:"type"`
	Path string `json:"path"`
	Fun  func(c *gin.Context)
}

func init() {
	api.Db.AutoMigrate(&api.User{})
	api.Db.AutoMigrate(&api.Chat{})
	api.Db.AutoMigrate(&api.Message{})
}

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	chatService := service.NewChatService()
	chatHandler := handler.NewChatHandler(chatService)

	//生成一个request的列表
	requests := []request{
		///TODO: User操作
		{Type: api.POST, Path: api.REGISTER, Fun: userHandler.RegisterUser},
		{Type: api.POST, Path: api.LOGIN, Fun: userHandler.LoginUser},
		{Type: api.POST, Path: api.FIND, Fun: userHandler.FindUser},
		{Type: api.GET, Path: api.FINDALL, Fun: userHandler.FindAllUser},
		{Type: api.GET, Path: api.DELETEUSER, Fun: userHandler.DeleteUser},

		///TODO: Chat操作
		{Type: api.POST, Path: api.StartAChatHAT, Fun: chatHandler.StartAChat},
		{Type: api.GET, Path: api.DELETECHAT, Fun: chatHandler.DeleteChat},
		{Type: api.GET, Path: api.DELETEALLCHAT, Fun: chatHandler.DeleteAllChat},
		{Type: api.GET, Path: api.GETCHATDETAIL, Fun: chatHandler.GetChatDetail},
		{Type: api.GET, Path: api.GETCHATLIST, Fun: chatHandler.GetChatList},
		{Type: api.POST, Path: api.SENDMESSAGE, Fun: chatHandler.SendMessage},
	}

	//注册路由
	r := gin.Default()

	//配置跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Next()
	})

	for _, v := range requests {
		path := api.API + v.Path
		if v.Type == api.POST {
			r.POST(path, v.Fun)
		} else if v.Type == api.GET {
			r.GET(path, v.Fun)
		} else if v.Type == api.DELETE {
			r.DELETE(path, v.Fun)
		} else {
			panic("error request type")
		}
	}

	r.GET(api.API+"/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ws", ws.Handler)

	r.Run(":8080")
}
