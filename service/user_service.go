package service

import (
	"github.com/habibiiberahim/go-backend/model"
)

type UserService interface {
	//
	FindAll() (response []model.GetProductResponse)
}
