package api

const (
	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
)

// url前缀
const (
	API = ""
)

// / user
const (
	REGISTER   = "/user/register"
	LOGIN      = "/user/login"
	FIND       = "/user/find"
	FINDALL    = "/user/findall"
	DELETEUSER = "/user/delete"
)

// / chat
const (
	StartAChatHAT = "/chat/start"
	DELETECHAT    = "/chat/delete"
	DELETEALLCHAT = "/chat/deleteall"
	GETCHATDETAIL = "/chat/detail"
	GETCHATLIST   = "/chat/list"
	SENDMESSAGE   = "/chat/send"
)
