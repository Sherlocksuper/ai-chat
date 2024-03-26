package ws

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
	"time"
)

var (
	// 消息通道
	news = make(map[string]chan interface{})
	// websocket客户端链接池
	client = make(map[string]*websocket.Conn)
	// 互斥锁，防止程序对统一资源同时进行读写
	mux sync.Mutex
)

// websocket
var upgrade = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		// 取消ws跨域校验
		return true
	},
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		log.Error().Str("reason", reason.Error()).Str("location", "websocket.go").Msg("websocket连接失败")
	},
}

// Handler 处理ws请求
func Handler(context *gin.Context) {
	userId := context.Query("userId")
	connect, err := upgrade.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		return
	}
	log.Info().Str("userId", userId).Msg("websocket连接成功")

	//把客户端添加到客户端链接池
	addClient(userId, connect)

	for {
		messageType, message, err := connect.ReadMessage()
		fmt.Println(messageType)
		if err != nil {
			log.Error().Str("userId", userId).Msg("websocket连接失败")
			deleteClient(userId)
			connect.Close()
			break
		}

		response := "向客户端发送消息：" + string(message)
		context.Writer.Write([]byte(response))

		connect.WriteMessage(messageType, []byte(response))
	}

	//关闭连接
	defer connect.Close()
}

// 将客户端添加到客户端链接池
func addClient(id string, conn *websocket.Conn) {
	mux.Lock()
	client[id] = conn
	mux.Unlock()
}

// 获取指定客户端链接
func getClient(id string) (conn *websocket.Conn, exist bool) {
	mux.Lock()
	conn, exist = client[id]
	mux.Unlock()
	return conn, exist
}

// 删除客户端链接
func deleteClient(id string) {
	mux.Lock()
	delete(client, id)
	log.Info().Msg("删除了客户端链接userId:" + id + "    location:websocket.go")
	mux.Unlock()
}

// SendMsg 发送消息
func SendMsg(urId int, message WsReMessage) {

	userId := fmt.Sprintf("%d", urId)
	connect, isExist := getClient(userId)

	if !isExist {
		return
	}

	//把message变为json字符串
	messageString, _ := json.Marshal(message)

	if message.Type != CHAT_MESSAGE {
		log.Info().Msg("向客户端发送消息：" + string(messageString) + "location:websocket.go")
	}

	connect.WriteMessage(websocket.TextMessage, messageString)
}
