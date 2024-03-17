package api

func M(code int, message string, data any) ReturnMessage {
	return ReturnMessage{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
