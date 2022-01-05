package service

import (
	"github.com/habibiiberahim/go-backend/model"
)

type UserService interface {
	//Function create membutuhkan request model lalu akan dikembalikan menjadi request response

	List() (response []model.GetProductResponse)
}
