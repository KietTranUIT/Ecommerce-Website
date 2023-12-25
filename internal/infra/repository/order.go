package repository

import (
	"log"
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

func (repo userRepository) GetOrderWithEmail(user_email string) []dto.Order {
	var orders []dto.Order
	result := repo.db.GetDB().Table("orders").
		Select("id as id, user_email as user_email, status as status, total as total, created_at as created_at").
		Where("user_email = ?", user_email).Scan(&orders)

	if result.Error != nil {
		var empty []dto.Order
		log.Println(result.Error.Error())
		return empty
	}
	return orders
}
