package authenticaion

import (
	"backend/config"
	templateError "backend/error"
	"errors"

	"gorm.io/gorm"
)

type AuthenRepository struct{}

func (ar *AuthenRepository) InitAuthenRepository() *AuthenRepository {
	return &AuthenRepository{}
}

func (ar *AuthenRepository) GetUser(email string) (userPassword *User, err error) {
	if config.DataBase.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.DataBase.DB
	if err = db.Table("users").Select("*").Where("email = ?", email).First(&userPassword).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.UsernotfoundError
		} else {
			return nil, err
		}
	}
	return
}

func (ar *AuthenRepository) GetUserByUID(uid string) (userPassword *User, err error) {
	if config.DataBase.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := config.DataBase.DB
	if err = db.Table("users").Select("*").Where("uid = ?", uid).First(&userPassword).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.UsernotfoundError
		} else {
			return nil, err
		}
	}
	return
}

func (ar *AuthenRepository) CreateUser(userData User) (err error) {
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
