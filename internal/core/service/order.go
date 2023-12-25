package service

import (
	"time"
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

func (service userService) CreateOrder(req request.CreateOrderRequest) *response.Response {
	var status string
	if req.Payment_id == 4 {
		status = "Thanh toan khi nhan hang"
	} else {
		status = "Da thanh toan"
	}

	now := time.Now()
	order := &dto.Order{
		User_email: req.User_email,
		Address_id: req.Address_id,
		Payment_id: req.Payment_id,
		Total:      req.Subtotal,
		Status:     status,
		Created_at: &now,
	}

	if result := service.repo.CreateOrder(order); result != nil {
		return CreateFailResponse(error_code.CreateOrderFail, result.Error())
	}

	var order_details []dto.OrderDetail

	for _, p := range req.Cart {
		detail := dto.OrderDetail{
			Order_id:   order.Id,
			Product_id: p.Product_id,
			Quantity:   p.Quantity,
		}
		order_details = append(order_details, detail)
	}

	if result := service.repo.CreateOrderDetails(order_details); result != nil {
		return CreateFailResponse(error_code.CreateOrderFail, result.Error())
	}

	return CreateSuccessResponse(error_code.Success, "")
}

func (service userService) GetOrderWithEmail(user_email string) []dto.Order {
	return service.repo.GetOrderWithEmail(user_email)
}
