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

func (service userService) DeleteUserAddress(req request.DeleteUserAddressRequest) *response.Response {
	result := service.repo.DeleteUserAddress(req.Id)
	if result != nil {
		return CreateFailResponse(error_code.DeleteUserAddressFail, result.Error())
	}
	return CreateSuccessResponse(error_code.Success, "")
}

func (service userService) EditUserAddress(req request.EditUserAddressRequest) *response.Response {
	data := &dto.UserAddress{
		Id:      req.Id,
		Address: req.Address,
		City:    req.City,
		Phone:   req.Phone,
	}

	result := service.repo.UpdateUserAddress(data)
	if result != nil {
		return CreateFailResponse(error_code.UpdateUserAddressFail, result.Error())
	}
	return CreateSuccessResponse(error_code.Success, "")
}

func (service userService) GetUserAddress(user_email string) ([]dto.UserAddress, error) {
	return service.repo.GetUserAddress(user_email)
}
