package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/service"
	"github.com/gin-gonic/gin"
)

type PromptHandler struct {
	promptService service.PromptService
}

func NewPromptHandler(service.PromptService) *PromptHandler {
	return &PromptHandler{
		promptService: service.NewPromptService(),
	}
}

// AddPrompt 添加一个提示 POST
func (p *PromptHandler) AddPrompt(c *gin.Context) {
	var prompt api.Prompt

	if err := c.BindJSON(&prompt); err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}

	if err := p.promptService.AddPrompt(&prompt); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "创建成功"})
}

// DeletePrompt 删除一个提示 GET  promptId
func (p *PromptHandler) DeletePrompt(c *gin.Context) {
	id := c.Query("promptId")

	if err := p.promptService.DeletePrompt(id); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "删除成功"})
}

// GetPromptList 获取所有提示 GET
func (p *PromptHandler) GetPromptList(c *gin.Context) {
	var prompts []api.Prompt

	if err := p.promptService.GetPromptList(&prompts); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "查找成功", "data": prompts})
}

// UpdatePrompt 更新一个提示 POST
func (p *PromptHandler) UpdatePrompt(c *gin.Context) {
	var prompt api.Prompt

	if err := c.BindJSON(&prompt); err != nil {
		c.JSON(400, gin.H{"message": "参数错误"})
		return
	}

	if err := p.promptService.UpdatePrompt(&prompt); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "更新成功"})
}
