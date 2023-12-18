package service

import (
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

func (service userService) CreateUserAddress(req request.CreateUserAddressRequest) *response.Response {
	data := &dto.UserAddress{
		User_email: req.User_email,
		Address:    req.Address,
		City:       req.City,
		Phone:      req.Phone,
	}

	result := service.repo.CreateUserAddress(data)
	if result != nil {
		return CreateFailResponse(error_code.CreateUserAddressFail, result.Error())
	}
	return CreateSuccessResponse(error_code.Success, "")
}
