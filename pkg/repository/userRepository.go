package repository

import (
	"athun/pkg/domain"
	interfaces "athun/pkg/repository/interface"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &userDatabase{
		DB: db,
	}
}

func (r *userDatabase) Create(user domain.User) error {
	result := r.DB.Create(&user).Error
	return result
}
