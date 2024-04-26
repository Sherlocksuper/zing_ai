package service

import (
	"awesomeProject3/api"
	api2 "awesomeProject3/api/struct"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService interface {
	Login(user *api.User) error
	HasNameOrEmail(name string, email string) error
	Register(username, email, password string) error
	ResetPassword(email string, password string) error
	Backup() error
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (auth authService) HasNameOrEmail(name string, email string) error {
	var user api.User

	api.Db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return errors.New("用户名已存在")
	}

	api.Db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return errors.New("邮箱已存在")
	}

	fmt.Println(user)

	return nil
}

func (auth authService) Register(username, email, password string) error {
	log.Info().Msg("用户名字:" + username + "密码:" + password + "  注册" + "location is :service/user.go  Register")
	err := auth.HasNameOrEmail(username, email)
	if err != nil {
		return err
	}

	//密码加密
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	api.Db.Create(&api.User{Name: username, Email: email, Password: string(fromPassword)})
	return nil
}

func (auth authService) ResetPassword(email string, password string) error {
	var user api.User
	api.Db.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	api.Db.Model(&user).Update("password", string(fromPassword))
	return nil
}

func (auth authService) Login(user *api.User) error {
	log.Info().Msg("用户" + user.Email + "密码" + user.Password + "   登录")
	password := user.Password
	api.Db.Where("email = ?", user.Email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("密码错误")
	}

	if user.AccountStatus == api2.Lock {
		return errors.New("账号被封禁")
	}

	now := time.Now()
	api.Db.Model(&user).Update("last_login_time", now.Format("2006-01-02 15:04:05"))

	return nil
}

func (auth authService) Backup() error {
	return nil
}
