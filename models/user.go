package models

type User struct {
	Base
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	UserName  string `json:"user_name" db:"user_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
}

type GetAllUsersResponse struct {
	Meta *Meta  `json:"meta"`
	Data []User `json:"data"`
}
