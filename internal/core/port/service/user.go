package service

import (
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

type UserService interface {
	SignUp(request.SignUpRequest) *response.Response
	SendVerificationCode(email string) *response.Response
	AuthenticateCode(request.AuthenticateRequest) *response.Response
}
