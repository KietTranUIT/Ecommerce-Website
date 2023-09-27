package repository

import (
	"user-service/internal/core/dto"
)

type UserRepository interface {
	CreateUser(user *dto.UserDTO) error
	CreateUserAddress(user_address *dto.UserAddress) error
	CreateVerificationMail(verify *dto.VerificationEmail) error
	GetVerificationWithEmailAndType(string, string) *dto.VerificationEmail
	UpdateVerificationEmail(dto.VerificationEmail) error
	UpdateStatusVerificationEmail(string, string) error
}
