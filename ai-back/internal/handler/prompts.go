package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
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
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
		return
	}

	if err := p.promptService.AddPrompt(&prompt); err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "添加成功", nil))
}

// DeletePrompt 删除一个提示 GET  promptId
func (p *PromptHandler) DeletePrompt(c *gin.Context) {
	id := c.Query("promptId")

	if err := p.promptService.DeletePrompt(id); err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "删除成功", nil))
}

// GetPromptList 获取所有提示 GET
func (p *PromptHandler) GetPromptList(c *gin.Context) {
	offset := c.Query("offset")
	fmt.Println("offset is :", offset, "location is :handler/prompts.go")

	intOffset, _ := strconv.Atoi(offset)

	var prompts []api.Prompt
	if err := p.promptService.GetPromptList(&prompts, intOffset); err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "获取成功", prompts))
}

// UpdatePrompt 更新一个提示 POST
func (p *PromptHandler) UpdatePrompt(c *gin.Context) {
	var prompt api.Prompt

	if err := c.BindJSON(&prompt); err != nil {
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
		return
	}

	if err := p.promptService.UpdatePrompt(&prompt); err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}

	c.JSON(200, api.M(api.SUCCESS, "更新成功", nil))
}
