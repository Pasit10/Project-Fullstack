package entities

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	UID      string `json:"uid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}
