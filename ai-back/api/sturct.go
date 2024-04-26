package api

import (
	_ "awesomeProject3/api/struct"
	api2 "awesomeProject3/api/struct"
	"github.com/gin-gonic/gin"
)

type ReturnMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Func struct {
	Url    string `json:"url"`
	Method string `json:"method"`
	Action func(c *gin.Context)
}

type User = api2.User
type Chat = api2.Chat
type Message = api2.Message
type Prompt = api2.Prompt
type Version = api2.Version
type Email = api2.Email
