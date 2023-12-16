// package contains interfaces to handle requests from user
package service

import (
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

type UserService interface {
	// Handle sign up
	SignUp(request.SignUpRequest) *response.Response

	Login(request.LoginRequest) *response.Response

	// Send a verification code to user email
	SendVerificationCode(email string) *response.Response

	// authenticate user with code
	AuthenticateCode(request.AuthenticateRequest) *response.Response

	CheckAccount(string) *response.Response
}
