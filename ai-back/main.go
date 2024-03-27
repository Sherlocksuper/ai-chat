package main

import (
	"awesomeProject3/docs"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)                                  // gin-swagger middleware
import _ "github.com/swaggo/files" // swagger embed files

import (
	"awesomeProject3/api"
	"awesomeProject3/handler"
	"awesomeProject3/service"
	"awesomeProject3/ws"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type request struct {
	Type string `json:"type"`
	Path string `json:"path"`
	Fun  func(c *gin.Context)
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	var err error

	err = api.Db.AutoMigrate(&api.User{})
	err = api.Db.AutoMigrate(&api.Chat{})
	err = api.Db.AutoMigrate(&api.Message{})
	err = api.Db.AutoMigrate(&api.Version{})
	err = api.Db.AutoMigrate(&api.Prompt{})
	if err != nil {
		log.Error().Msg("数据库迁移失败：" + err.Error())
	}
}

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	chatService := service.NewChatService()
	chatHandler := handler.NewChatHandler(chatService)

	versionService := service.NewVersionService()
	versionHandler := handler.NewVersionHandler(versionService)

	promptService := service.NewPromptService()
	promptHandler := handler.NewPromptHandler(promptService)

	r := gin.Default()

	userGroup := r.Group(api.API + "/user")
	{
		userGroup.POST("/register", userHandler.RegisterUser)
		userGroup.POST("/login", userHandler.LoginUser)
		userGroup.POST("/find", userHandler.FindUser)
		userGroup.GET("/findAll", userHandler.FindAllUser)
		userGroup.GET("/delete", userHandler.DeleteUser)
		userGroup.GET("/getemailcode", userHandler.GetEmailCode)
		userGroup.GET("/checkemailcode", userHandler.CheckRegisterCode)
	}

	chatGroup := r.Group(api.API + "/chat")
	{
		chatGroup.POST("/start", chatHandler.StartAChat)
		chatGroup.GET("/delete", chatHandler.DeleteChat)
		chatGroup.GET("/deleteall", chatHandler.DeleteAllChat)
		chatGroup.GET("/detail", chatHandler.GetChatDetail)
		chatGroup.GET("/list", chatHandler.GetChatList)
		chatGroup.POST("/send", chatHandler.SendMessage)
	}

	versionGroup := r.Group(api.API + "/version")
	{
		versionGroup.GET("/all", versionHandler.GetAllVersions)
		versionGroup.POST("/add", versionHandler.AddVersion)
		versionGroup.GET("/latest", versionHandler.GetLatestVersion)
	}

	promptGroup := r.Group(api.API + "/prompt")
	{
		promptGroup.POST("/add", promptHandler.AddPrompt)
		promptGroup.GET("/delete", promptHandler.DeletePrompt)
		promptGroup.GET("/list", promptHandler.GetPromptList)
		promptGroup.POST("/update", promptHandler.UpdatePrompt)
	}

	//配置跨域
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Next()
	})

	r.GET(api.API+"/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ws", ws.Handler)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
