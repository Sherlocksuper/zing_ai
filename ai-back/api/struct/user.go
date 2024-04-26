package api

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string        `json:"name"`
	Password      string        `json:"password"`
	Token         string        `json:"token"`
	Email         string        `json:"email"`
	Chats         []Chat        `json:"chats"`
	LastLoginTime string        `json:"lastLoginTime"`
	ChatNum       int           `json:"chatNum" gorm:"default:0"`
	Role          Role          `json:"role" gorm:"default:User"`
	AccountStatus AccountStatus `json:"accountStatus"`
}

type Role string

const (
	Customer Role = "User"
	Admin    Role = "Admin"
)

type AccountStatus string

const (
	Normal AccountStatus = "Normal"
	Lock   AccountStatus = "Lock"
)

// AfterCreate 创建用户后,如果用户id为1,则为管理员
func (u *User) AfterCreate(tx *gorm.DB) error {
	role := Customer
	if u.ID == 1 {
		role = Admin
	}
	tx.Model(&User{}).Where("id = ?", u.ID).Update("role", role)
	return nil
}
