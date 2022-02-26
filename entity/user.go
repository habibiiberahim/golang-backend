package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	FullName string `gorm:"size:255"`
	Email    string `gorm:"size:100:not null:unique"`
	SocialId string `gorm:"size:100:not null:unique"`
	Provider string `gorm:"size:255"`
	Password string `gorm:"size:255:default:''"`
}
