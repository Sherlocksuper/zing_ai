package api

import "gorm.io/gorm"

type Prompt struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Function string `json:"function"`
}
