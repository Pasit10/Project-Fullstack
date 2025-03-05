package repositories

import (
	templateError "backend/error"
	"backend/internal/entities"
	"errors"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

type IAuthRepository interface {
	GetUser(email string) (userPassword *entities.User, err error)
	GetUserByUID(uid string) (userPassword *entities.User, err error)
	CreateUser(userData entities.User) (err error)
}

func InitAuthRepository(db *gorm.DB) IAuthRepository {
	return authRepository{
		DB: db,
	}
}

func (repo authRepository) GetUser(email string) (userPassword *entities.User, err error) {
	if repo.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	if err = repo.DB.Table("users").Select("*").Where("email = ?", email).First(&userPassword).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.UsernotfoundError
		} else {
			return nil, err
		}
	}
	return
}

func (repo authRepository) GetUserByUID(uid string) (userPassword *entities.User, err error) {
	if repo.DB == nil {
		return nil, templateError.DatabaseConnectedError
	}
	db := repo.DB
	if err = db.Table("users").Select("*").Where("uid = ?", uid).First(&userPassword).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, templateError.UsernotfoundError
		} else {
			return nil, err
		}
	}
	return
}

func (repo authRepository) CreateUser(userData entities.User) (err error) {
	if repo.DB == nil {
		return nil
	}
	err = repo.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("users").Create(&userData).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	return
}
