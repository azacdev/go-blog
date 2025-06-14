package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"varchar:191"`
	Email        string `gorm:"varchar:191;unique"`
	Password     string `gorm:"varchar:191"`
	Picture      string `gorm:"varchar:255"`
	RefreshToken string `gorm:"varchar:255"`
}
