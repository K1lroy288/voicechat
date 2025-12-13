package repository

import (
	"user-service/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Register(user *model.User) (bool, error) {
	var exist model.User
	res := r.DB.Where("username = ?", user.Username).First(&exist).Error
	if res != nil {
		err := r.DB.Create(user).Error
		return false, err
	}
	return res == nil, nil
}

func (r *UserRepository) GetUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *UserRepository) GetUserById(id int) (model.User, error) {
	var user model.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
