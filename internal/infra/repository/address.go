package repository

import (
	"user-service/internal/core/dto"
)

func (repo userRepository) CreateUserAddress(userAddress *dto.UserAddress) error {
	result := repo.db.GetDB().Table("user_address").Select("user_email, address, city, phone").Create(&userAddress)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
