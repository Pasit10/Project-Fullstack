package authenticaion

import (
	templateError "backend/error"
	"backend/utils"
	"errors"
	"fmt"
	"net/mail"
)

type AuthenLogic struct {
	AuthenRepository *AuthenRepository
}

func (al *AuthenLogic) InitAuthenLogic() *AuthenLogic {
	repo := &AuthenRepository{}
	return &AuthenLogic{
		AuthenRepository: repo.InitAuthenRepository(),
	}
}

func (al *AuthenLogic) LoginLogic(userData UserLogin) (isValid bool, user *User, err error) {
	if userData.Email == "" {
		return false, nil, templateError.BadrequestError
	}
	if _, err := mail.ParseAddress(userData.Email); err != nil {
		return false, nil, templateError.BadrequestError
	}
	if userData.Password == "" {
		return false, nil, templateError.BadrequestError
	}
	user, err = al.AuthenRepository.GetUser(userData.Email)
	if err != nil {
		if errors.Is(err, templateError.UsernotfoundError) {
			return false, nil, templateError.WrongUserOrPasswordError
		} else {
			fmt.Println(err)
			return false, nil, err
		}
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

func (al *AuthenLogic) RegisterLogic(userData User) (err error) {
	if userData.Email == "" {
		return templateError.BadrequestError
	}
	if _, err := mail.ParseAddress(userData.Email); err != nil {
		return errors.New("invalid email format")
	}
	if userData.Password == "" {
		return templateError.BadrequestError
	}
	_, err = al.AuthenRepository.GetUser(userData.Email)
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

	err = al.AuthenRepository.CreateUser(userData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func (al *AuthenLogic) RegisterGoogleLogic(userData User) (err error) {
	if userData.Email == "" {
		return templateError.BadrequestError
	}
	if _, err := mail.ParseAddress(userData.Email); err != nil {
		return errors.New("invalid email format")
	}
	_, err = al.AuthenRepository.GetUser(userData.Email)
	if !errors.Is(err, templateError.UsernotfoundError) {
		return templateError.EmailAlreadyExistError
	}

	err = al.AuthenRepository.CreateUser(userData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return
}

func (al *AuthenLogic) GetUserByUID(uid string) (user *User, err error) {
	user, err = al.AuthenRepository.GetUserByUID(uid)
	if err != nil {
		return nil, err
	}
	return
}
