package database

import (
	"github.com/antoniomjr/go/9-apis/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	//FindAll(page, limit int, sort string) ([]entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	//Delete(email string) error
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
