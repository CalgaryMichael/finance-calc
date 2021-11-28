package models

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateUserRequest struct {
	User User `json:"user"`
}

type CreateUserResponse struct {
	UserId int `json:"userId"`
}
