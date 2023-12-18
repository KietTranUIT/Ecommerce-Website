// Package contains interfaces comunicate with Database
package repository

import (
	"user-service/internal/core/dto"
)

type UserRepository interface {
	// Insert a user into Users table in database
	CreateUser(user *dto.UserDTO) error

	// Insert a user address into Users_address table in database
	CreateUserAddress(user_address *dto.UserAddress) error

	// Insert a verification into Email_verification table in database
	CreateVerificationMail(verify *dto.VerificationEmail) error

	// Select a verification from Email_verification table with email and type conditions
	GetVerificationWithEmailAndType(string, string) *dto.VerificationEmail

	// Update (code, expire time) for a verification
	UpdateVerificationEmail(dto.VerificationEmail) error

	// Update (status) for a verification
	UpdateStatusVerificationEmail(string, string) error

	GetUserWithEmail(string) *dto.UserDTO

	// For Admin -----------------------------------
	GetAdmin(string) *dto.AdminDTO
	GetProductCategories() []dto.ProductCategory
	GetLastIDCategories() int
	InsertCategory(dto.ProductCategory) error
	DeleteCategory(id int) error
	GetCategoryWithId(id string) *dto.ProductCategory
	UpdateCategory(*dto.ProductCategory) error

	GetProductForAdmin() []dto.Product
	InsertProduct(*dto.Product) error
	GetProductWithId(string) *dto.Product
	UpdateProduct(*dto.Product) error
	GetProductVersion(string) []dto.ProductVersion
	GetLastIdProductVersion() int
	CreateProductVersion(*dto.ProductVersion) error
	CreateProductInventory(*dto.ProductInventory) error
	UpdateProductVersion(*dto.ProductVersion) error
	UpdateProductInventory(int, int) error
	GetProductVersionWithId(int) *dto.ProductVersion
	DeleteProductInventory(int) error
	DeleteProductVersion(int) error
	GetOrderAdminPage() []dto.Order
	GetOrderDetail(int) []dto.OrderDetail
	GetOrderWithId(int) *dto.Order
}
