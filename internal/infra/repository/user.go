package repository

import (
	"errors"
	"user-service/internal/core/dto"
	"user-service/internal/core/port/repository"
)

const rowAffected = 1

var (
	rowAffectedError = errors.New("row affected error")
)

type userRepository struct {
	db repository.Database
}

func NewUserRepository(db repository.Database) repository.UserRepository {
	return userRepository{
		db: db,
	}
}

func (repo userRepository) GetUserWithEmail(email string) *dto.UserDTO {
	var result dto.UserDTO
	repo.db.GetDB().Table("Users").Select("email", "password").Where("email = ?", email).Scan(&result)
	if result.Id == 0 && result.Email == "" {
		return nil
	}
	return &result
}

func (repo userRepository) CreateUser(user *dto.UserDTO) error {
	result := repo.db.GetDB().Create(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return rowAffectedError
	}

	return nil
}

func (repo userRepository) CreateUserAddress(user_address *dto.UserAddress) error {
	result := repo.db.GetDB().Create(user_address)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return rowAffectedError
	}

	return nil
}

func (repo userRepository) CreateVerificationMail(verify *dto.VerificationEmail) error {
	result := repo.db.GetDB().Create(verify)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return rowAffectedError
	}

	return nil
}

func (repo userRepository) GetVerificationWithEmailAndType(email string, type_verify string) *dto.VerificationEmail {
	var result dto.VerificationEmail
	repo.db.GetDB().Table("Email_verification").Where("email = ? AND type = ?", email, type_verify).First(&result)

	if result.Id == 0 {
		return nil
	}

	return &result

}

func (repo userRepository) UpdateVerificationEmail(verify dto.VerificationEmail) error {
	result := repo.db.GetDB().Table("Email_verification").Where("email = ? AND type = ?", verify.Email, verify.Type).Updates(map[string]interface{}{
		"code":      verify.Code,
		"expire_at": verify.Expire_at,
	})
	return result.Error
}

func (repo userRepository) UpdateStatusVerificationEmail(email string, kind string) error {
	result := repo.db.GetDB().Table("Email_verification").Where("email = ? AND type = ?", email, kind).Update("status", true)
	return result.Error
}
