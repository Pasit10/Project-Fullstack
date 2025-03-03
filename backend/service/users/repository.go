package users

import (
	"backend/config"
	templateError "backend/error"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct{}

func (ur *userRepository) InitUserRepository() *userRepository {
	return &userRepository{}
}

func (ur *userRepository) CreateUser(userData User) (err error) {
	if config.DataBase.DB == nil {
		return nil
	}
	db := config.DataBase.DB
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("users").Create(&userData).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}

func (ur *userRepository) GetUserByUID(uid string) (user *User, err error) {
	if config.DataBase.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.DataBase.DB

	if err = db.Table("users").Select("password").Where("uid = ?", uid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.UsernotfoundError
		} else {
			return nil, err
		}
	}
	return
}

func (ur *userRepository) GetUserByEmail(uid string) (user *User, err error) {
	if config.DataBase.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.DataBase.DB

	if err = db.Table("users").Select("*").Where("uid = ?", uid).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.UsernotfoundError
		} else {
			return nil, err
		}
	}
	return
}
