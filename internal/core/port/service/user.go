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
	//GetCategories() []dto.ProductCategory
	GetLastIDCategories() int
	DeleteCategory(id int) *response.Response
	GetCategoryWithId(id string) *dto.ProductCategory

	// Products Admin
	CreateProduct(*dto.Product) bool
	GetProductsForAdmin() []dto.Product
	GetProductWithId(string) *dto.Product
	UpdateProduct(*dto.Product) bool
	GetProductVersion(string) []dto.ProductVersion
	GetLastIdProduct() int
	DeleteProduct(int) *response.Response
	CreateProductVersion(*dto.ProductVersion) bool
	CreateProductInventory(*dto.ProductInventory) bool
	UpdateProductInventory(int, int) bool
	GetLastIdProductVersion() int
	UpdateProductVersion(*dto.ProductVersion) bool
	GetProductVersionWithId(int) *dto.ProductVersion
	DeleteProductVersion(int) bool
	GetOrderAdminPage() []dto.Order
	GetOrderDetail(int) []dto.OrderDetail
	GetOrderWithId(int) *dto.Order

	// Address user
	CreateUserAddress(request.CreateUserAddressRequest) *response.Response
	DeleteUserAddress(request.DeleteUserAddressRequest) *response.Response
	EditUserAddress(request.EditUserAddressRequest) *response.Response
	CreateOrder(request.CreateOrderRequest) *response.Response
	GetProductForHomePage() *response.Response
	GetProductData(int) []dto.ProductDTO

	// Category
	GetCategories() ([]dto.Category, error)
	CreateCategory(request.CreateCategoryRequest) *response.Response
	GetCategory(int) (*dto.Category, error)
	UpdateCategory(*dto.Category) *response.Response

	// Product
	GetProductsWithCategoryId(id int) *response.Response

	GetTotalSalesDayNow() *response.Response
	GetTotalSalesWeekNow() *response.Response
	GetTotalSalesMonthNow() *response.Response
	GetOrdersRecently() *response.Response
	GetTopProducts() *response.Response
	GetProductsWithCategoryIdV1(id int) []dto.Product
}
