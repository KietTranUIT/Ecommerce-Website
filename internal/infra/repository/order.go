package repository

import (
	"user-service/internal/core/dto"
)

func (repo userRepository) CreateOrder(order *dto.Order) error {
	result := repo.db.GetDB().Table("orders").
		Select("user_email", "address_id", "payment_id", "status", "total").
		Create(&order)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) CreateOrderDetails(orders []dto.OrderDetail) error {
	number := len(orders)
	result := repo.db.GetDB().Table("order_detail").
		Select("order_id", "product_id", "quantity").
		CreateInBatches(&orders, number)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
