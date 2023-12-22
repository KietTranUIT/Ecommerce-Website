package service

import (
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/response"
)

func (service userService) GetProductForHomePage() *response.Response {
	products, err := service.repo.GetProductForHomePage()

	if err != nil {
		return CreateFailResponse(error_code.GetAllProductsFail, err.Error())
	}
	return CreateSuccessResponse(error_code.Success, "", products)
}

func (service userService) GetProductData(id int) []dto.ProductDTO {
	return service.repo.GetProductData(id)
}

func (service userService) GetProductsWithCategoryId(id int) *response.Response {
	products, err := service.repo.GetProductsWithCategoryId(id)

	if err != nil {
		return CreateFailResponse(error_code.GetProductsWithCategoryId, err.Error())
	}
	return CreateSuccessResponse(error_code.Success, "", products)
}

func (service userService) GetProductsWithCategoryIdV1(id int) []dto.Product {
	products, err := service.repo.GetProductsWithCategoryId(id)

	if err != nil {
		return nil
	}
	return products
}

func (service userService) CreateProduct(product *dto.Product) bool {
	result := service.repo.InsertProduct(product)

	if result != nil {
		return false
	}
	return true
}
