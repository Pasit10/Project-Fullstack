package users

type userLogic struct {
	UserRepository *userRepository
}

func (ul *userLogic) InitUserLogic() *userLogic {
	repo := &userRepository{}
	return &userLogic{
		UserRepository: repo.InitUserRepository(),
	}
}

func (ul *userLogic) CreateUserLogic(user User) error {
	user.Role = "user"
	if err := ul.UserRepository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (ul *userLogic) GetuserData(uid string) (user *User, err error) {
	user, err = ul.UserRepository.GetUserByEmail(uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}
