package entity

import "gorm.io/gorm"

type User struct {
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
	gorm.Model
}
