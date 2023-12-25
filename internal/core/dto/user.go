package dto

import (
	"time"
)

type Tabler interface {
	TableName() string
}

type UserDTO struct {
	Id          uint
	Email       string
	Password    string
	First_name  string
	Last_name   string
	Gender      string
	Created_at  time.Time
	Modified_at time.Time
	Active      bool
}

func (U *UserDTO) TableName() string {
	return "Users"
}

func (U *UserAddress) TableName() string {
	return "User_address"
}

type VerificationEmail struct {
	Id        uint
	Email     string
	Code      string
	Status    bool
	Type      string
	Expire_at time.Time
}

type DataSales struct {
	Sales   int `json:"sales"`
	Revenue int `json:"revenue"`
}

func (v *VerificationEmail) TableName() string {
	return "Email_verification"
}

type PaymentMethod struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
