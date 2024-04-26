// user.go

package service

import (
	"awesomeProject3/api"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	HasNameOrEmail(name string, email string) error

	Register(username, email, password string) error
	Login(user *api.User) error
	ResetPassword(email uint, password string) error

	FindUser(user *api.User) error
	FindAllUser(users *[]api.User) error
	UpdateUser(user *api.User) error
}

type userService struct{}

func (u userService) HasNameOrEmail(name string, email string) error {
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

func (u userService) Register(username, email, password string) error {
	log.Info().Msg("用户名字:" + username + "密码:" + password + "  注册" + "location is :service/user.go  Register")
	err := u.HasNameOrEmail(username, email)
	if err != nil {
		return err
	}

	//密码加密
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	api.Db.Create(&api.User{Name: username, Email: email, Password: string(fromPassword)})
	return nil
}

func (u userService) ResetPassword(id uint, password string) error {
	var user api.User
	api.Db.First(&user, id)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	fromPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	api.Db.Model(&user).Update("password", string(fromPassword))
	return nil
}

func (u userService) Login(user *api.User) error {
	log.Info().Msg("用户" + user.Email + "密码" + user.Password + "   登录")
	password := user.Password
	api.Db.Where("email = ?", user.Email).First(&user)
	if user.ID == 0 {
		return errors.New("用户不存在")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("密码错误")
	}
	return nil
}

func (u userService) FindUser(user *api.User) error {
	// 通过id查找用户
	err := api.Db.Model(&api.User{}).Preload("Chats").Find(&user)
	//如果找不到user
	if err.RowsAffected == 0 {
		return errors.New("用户不存在")
	}
	return nil
}

func (u userService) FindAllUser(users *[]api.User) error {
	api.Db.Find(&users)
	return nil
}

// UpdateUser  更新用户信息
func (u userService) UpdateUser(user *api.User) error {
	userMap := map[string]interface{}{
		"id":              user.ID,
		"updated_at":      user.UpdatedAt,
		"name":            user.Name,
		"password":        user.Password,
		"token":           user.Token,
		"email":           user.Email,
		"account_status":  user.AccountStatus,
		"chat_num":        user.ChatNum,
		"role":            user.Role,
		"last_login_time": user.LastLoginTime,
	}

	log.Info().Msg("用户" + user.Email + "更新信息")

	for k, v := range userMap {
		if v == 0 || v == "" || len(v.([]byte)) == 0 {
			delete(userMap, k)
		}
	}

	if user.Role != "User" && user.Role != "Admin" {
		delete(userMap, "role")
	}

	if user.AccountStatus != "normal" && user.AccountStatus != "ban" {
		delete(userMap, "account_status")
	}

	err := api.Db.Save(&userMap)

	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewUserService() UserService {
	return &userService{}
}
