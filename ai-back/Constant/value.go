package Constant

var DefaultSystemMessage = "你是我的AI助手"

func JudgeEmpty(str string) bool {
	return len(str) == 0 || str == ""
}
