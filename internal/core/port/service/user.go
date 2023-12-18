// package contains interfaces to handle requests from user
package service

import (
	"user-service/internal/core/dto"
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

	// Interface for Admin
	LoginAdmin(request.LoginRequest) *response.Response
	GetCategories() []dto.ProductCategory
	GetLastIDCategories() int
	CreateCategory(dto.ProductCategory) bool
	DeleteCategory(id int) bool
	GetCategoryWithId(id string) *dto.ProductCategory
	UpdateCategory(*dto.ProductCategory) bool

	// Products Admin
	CreateProduct(*dto.Product) bool
	GetProductsForAdmin() []dto.Product
	GetProductWithId(string) *dto.Product
	UpdateProduct(*dto.Product) bool
	GetProductVersion(string) []dto.ProductVersion
	GetLastIdProductVersion() int
	CreateProductVersion(*dto.ProductVersion) bool
	CreateProductInventory(*dto.ProductInventory) bool
	UpdateProductInventory(int, int) bool
	UpdateProductVersion(*dto.ProductVersion) bool
	GetProductVersionWithId(int) *dto.ProductVersion
	DeleteProductVersion(int) bool
	GetOrderAdminPage() []dto.Order
	GetOrderDetail(int) []dto.OrderDetail
	GetOrderWithId(int) *dto.Order

	// Address user
	CreateUserAddress(request.CreateUserAddressRequest) *response.Response
}
