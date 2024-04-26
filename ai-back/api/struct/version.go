package api

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	Version      string `json:"version" gorm:"unique"`
	Introduction string `json:"introduction" gorm:"default:无描述"`
	Enable       bool   `json:"enable"`
	DownloadUrl  string `json:"downloadUrl"`
}
