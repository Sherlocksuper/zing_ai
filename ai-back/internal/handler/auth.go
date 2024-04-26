package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register 注册接口
func (auth *AuthHandler) Register(c *gin.Context) {
	var user api.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = auth.authService.Register(user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "注册成功", nil))
}

// ResetPassword 重置密码
func (auth *AuthHandler) ResetPassword(c *gin.Context) {
	var user api.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = auth.authService.ResetPassword(user.Email, user.Password)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "修改成功", nil))
}

// Login 登录接口
func (auth *AuthHandler) Login(c *gin.Context) {
	log.Debug().Msg("login")
	var user api.User
	err := c.BindJSON(&user)
	if err != nil || user.Email == "" || user.Password == "" {
		c.JSON(200, api.M(api.FAIL, "邮箱或密码为空", nil))
		return
	}
	err = auth.authService.Login(&user)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "登录成功", user))
}
