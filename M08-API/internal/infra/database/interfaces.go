package database

import "github.com/amauryeuzebio/goexpert/M08-API/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}
