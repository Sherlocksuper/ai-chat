package api

const (
	SUCCESS     = 200
	FAIL        = 400
	NOTFOUND    = 404
	SERVERERROR = 500
)

const (
	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
)

// /TODO url前缀
const (
	API = "aiapi"
)

// /TODO user
const (
	REGISTER       = "/user/register"
	LOGIN          = "/user/login"
	FIND           = "/user/find"
	FINDALL        = "/user/findall"
	DELETEUSER     = "/user/delete"
	GETEMAILCODE   = "/user/getemailcode"
	CHECKEMAILCODE = "/user/checkemailcode"
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

// /version
const (
	ALLVERSION    = "/version/all"
	ADDVERSION    = "/version/add"
	LATESTVERSION = "/version/latest"
)
