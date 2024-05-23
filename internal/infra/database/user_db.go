package database

import (
	"gorm.io/gorm"

	"github.com/antoniomjr/go/9-apis/internal/entity"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) Update(user *entity.User) error {
	_, err := u.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	return u.DB.Save(user).Error
}