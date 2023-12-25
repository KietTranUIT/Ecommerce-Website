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

func (service userService) GetProductData(id int) *dto.Item {
	product_version := service.repo.GetProductData(id)

	var product dto.Item
	product.Name = product_version[0].Name
	product.Image = product_version[0].Image
	product.Image1 = product_version[0].Image1
	product.Image2 = product_version[0].Image2
	product.Price = product_version[0].Price
	product.Description = product_version[0].Description
	product.Category_name = product_version[0].Category_name
	product.Category_id = product_version[0].Category_id

	for _, p := range product_version {
		product.Value = append(product.Value, dto.Item_value{Id: p.Id, Size_product: p.Size_product})
	}

	return &product
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
