package models

import "gorm.io/gorm"

type User struct {
	gorm.Model // 追加
	FirstName  string
	LastName   string
	Email      string `gorm:"unique"`
	Password   string
}
