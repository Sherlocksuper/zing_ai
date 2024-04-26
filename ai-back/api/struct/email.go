package api

import "gorm.io/gorm"

type Email struct {
	gorm.Model
	TargetEmail string `json:"targetEmail"`
	Title       string `json:"title"`
	Content     string `json:"content"`
}
