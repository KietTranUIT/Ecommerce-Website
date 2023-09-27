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

type UserAddress struct {
	Id        uint
	User_id   uint
	Telephone string
	Address   string
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

func (v *VerificationEmail) TableName() string {
	return "Email_verification"
}
