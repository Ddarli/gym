package models

type UserModel struct {
	Id          int    `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
	Password    string `json:"password" db:"password"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
}
