package dto

type UserAddress struct {
	Id         int    `json:"id"`
	User_email string `json:"user_email"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
}
