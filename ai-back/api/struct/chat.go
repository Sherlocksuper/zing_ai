package api

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"strconv"
)

type Chat struct {
	gorm.Model
	Title         string    `json:"title" gorm:"default:默认标题"`
	UserID        uint      `json:"userId"`
	SystemMessage string    `json:"systemMessage"`
	Messages      []Message `json:"messages"`
	Type          ChatType  `json:"type" gorm:"default:Text"`
}

type ChatType string

const (
	Text  ChatType = "Text"
	Image ChatType = "Image"
)

func (c *Chat) AfterCreate(tx *gorm.DB) error {
	// 获取用户聊太难数量
	var user User
	tx.Model(&User{}).First(&user, c.UserID)
	num := user.ChatNum
	// 更新用户聊天数量
	err := tx.Model(&User{}).Where("id = ?", c.UserID).Update("chat_num", gorm.Expr("chat_num + ?", num+1))

	log.Info().Strs("chat", []string{c.Title}).Str("userId", strconv.Itoa(int(c.UserID))).Msg("用户开始一个聊天")

	if err.Error != nil {
		log.Error().Msg("更新用户聊天数量失败")
		return err.Error
	}

	return nil
}
