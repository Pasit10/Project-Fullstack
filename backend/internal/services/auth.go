package services

import (
	templateError "backend/error"
	"backend/internal/entities"
	"backend/internal/repositories"
	"backend/pkg/utils"
	"errors"
	"fmt"
	"net/mail"
)

type authService struct {
	AuthRepository repositories.IAuthRepository
}

type IAuthService interface {
	Login(userData entities.UserLogin) (isValid bool, user *entities.User, err error)
	Register(userData entities.User) (err error)
	RegisterGoogle(userData entities.User) (err error)
	GetUserByUID(uid string) (user *entities.User, err error)
}

func InitAuthenService(repo repositories.IAuthRepository) IAuthService {
	return &authService{
		AuthRepository: repo,
	}
}

func (ser authService) Login(userData entities.UserLogin) (isValid bool, user *entities.User, err error) {
	if userData.Email == "" {
		return false, nil, templateError.BadrequestError
	}
	if _, err := mail.ParseAddress(userData.Email); err != nil {
		return false, nil, templateError.BadrequestError
	}
	if userData.Password == "" {
		return false, nil, templateError.BadrequestError
	}
	user, err = ser.AuthRepository.GetUser(userData.Email)
	if err != nil {
		if errors.Is(err, templateError.UsernotfoundError) {
			return false, nil, templateError.WrongUserOrPasswordError
		} else {
			fmt.Println(err)
			return false, nil, err
		}
	}
	if user.Password == "" {
		return false, nil, nil
	}
	isValid, err = utils.VerifyPassword(userData.Password, user.Password)
	if err != nil {
		fmt.Println(err)
		return false, nil, err
	}
	if !isValid {
		return false, nil, templateError.WrongUserOrPasswordError
	}
	return
}

func (ser authService) Register(userData entities.User) (err error) {
	if userData.Email == "" {
		return templateError.BadrequestError
	}
	if _, err := mail.ParseAddress(userData.Email); err != nil {
		return templateError.EmailInvaildFormatError
	}
	if userData.Password == "" {
		return templateError.BadrequestError
	}
	_, err = ser.AuthRepository.GetUser(userData.Email)
	if !errors.Is(err, templateError.UsernotfoundError) {
		return templateError.EmailAlreadyExistError
	}

	// hash password
	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	userData.Password = hashedPassword

	err = ser.AuthRepository.CreateUser(userData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func (ser authService) RegisterGoogle(userData entities.User) (err error) {
	if userData.Email == "" {
		return templateError.BadrequestError
	}
	if _, err := mail.ParseAddress(userData.Email); err != nil {
		return templateError.EmailInvaildFormatError
	}
	_, err = ser.AuthRepository.GetUser(userData.Email)
	if !errors.Is(err, templateError.UsernotfoundError) {
		return templateError.EmailAlreadyExistError
	}

	err = ser.AuthRepository.CreateUser(userData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func (ser authService) GetUserByUID(uid string) (user *entities.User, err error) {
	user, err = ser.AuthRepository.GetUserByUID(uid)
	if err != nil {
		return nil, err
	}
	return
}
