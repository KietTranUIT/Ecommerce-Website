package repository

import (
	"user-service/internal/core/dto"
)

func (repo userRepository) CreateUserAddress(userAddress *dto.UserAddress) error {
	result := repo.db.GetDB().Table("user_address").Select("user_email", "address", "city", "phone").Create(&userAddress)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) DeleteUserAddress(id int) error {
	result := repo.db.GetDB().Table("user_address").Where("id = ?", id).Delete(nil)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) UpdateUserAddress(userAddress *dto.UserAddress) error {
	result := repo.db.GetDB().Table("user_address").Where("id = ?", userAddress.Id).Updates(
		map[string]interface{}{
			"address": userAddress.Address,
			"city":    userAddress.City,
			"phone":   userAddress.Phone,
		})

	if result.Error != nil {
		return result.Error
	}
	return nil
}
