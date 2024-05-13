package database

import "github.com/antoniomjr/go/9-apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}