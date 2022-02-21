package model

import "gorm.io/gorm"

type GetProductResponse struct {
	gorm.Model
	Name  string
	Email string
}
