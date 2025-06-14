package models

import (
	"github.com/azacdev/go-blog/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"varchar: 191"`
	Content string `gorm:"text"`
	// Image   string `gorm:"text"`
	UserID uint
	User   models.User
}
