// handler.go

package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/config"
	service2 "awesomeProject3/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type UserHandler struct {
	userService service2.UserService
}

func NewUserHandler(userService service2.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetEmailCode GET  参数: email(api.Email)
func (f *UserHandler) GetEmailCode(c *gin.Context) {
	var err error
	var userEmail string
	userEmail = c.Query("email")

	redisService := service2.NewRedisService()
	emailService := service2.NewEmailService()

	//生成、redis储存code
	var code string
	for i := 0; i < 6; i++ {
		code += fmt.Sprintf("%d", rand.Intn(10))
	}
	fmt.Println("给"+userEmail+"的验证码为："+code, "   location is :handler/user.go  GetEmailCode")
	err = redisService.Set("registerCode", code)

	var content = fmt.Sprintf(config.GetConfig().Email.EmailTemplate, code)
	err = emailService.SendEmail(userEmail, config.GetConfig().Email.EmailTitle, content)

	if err != nil {
		c.JSON(200, api.M(api.FAIL, "发送失败", err.Error()))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "发送成功", nil))
}

// CheckRegisterCode 因为register要post入整个user，不太好融入，所以分开了
func (f *UserHandler) CheckRegisterCode(c *gin.Context) {
	//获取get参数，这里是email和code
	email := c.Query("email")
	code := c.Query("code")

	redisService := service2.NewRedisService()
	registerCode, _ := redisService.Get("registerCode")
	if code == registerCode {
		redisService.Set(email, "1")
		c.JSON(200, api.M(api.SUCCESS, "验证成功", nil))
	} else {
		c.JSON(200, api.M(api.FAIL, "验证失败", "验证码错误"))
	}
}

// FindUser 查找用户
func (f *UserHandler) FindUser(c *gin.Context) {
	var user api.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = f.userService.FindUser(&user)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "查找成功", user))
}

// FindAllUser GET 查找所有用户
func (f *UserHandler) FindAllUser(c *gin.Context) {
	var users []api.User
	err := f.userService.FindAllUser(&users)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "查找成功", users))
}

// UpdateUser 更新用户
func (f *UserHandler) UpdateUser(c *gin.Context) {
	var user api.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = f.userService.UpdateUser(&user)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "更新成功", nil))
}
