package api

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ChatID  int    `json:"chatId"`
	Role    string `json:"role"`
	Content string `json:"content"`
}
