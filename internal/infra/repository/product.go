package repository

import (
	"user-service/internal/core/dto"
)

func (repo userRepository) GetProductForHomePage() ([]dto.Product, error) {
	var products []dto.Product
	sql := `SELECT product.id as id, category.name as category_name, product.image as image, product.name as name, product.price as price FROM product inner join category on product.category_id = category.id`

	result := repo.db.GetDB().Raw(sql).Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo userRepository) GetProductData(id int) []dto.ProductDTO {
	var products []dto.ProductDTO
	repo.db.GetDB().Table("product").
		Select("product.name as name, product_version.id as id, category.name as category_name, category.id as category_id, product.description as description, product_version.size_product as size_product, product.price as price, product.image as image, product.image1 as image1, product.image2 as image2").
		Joins("inner join category on category.id = product.category_id").
		Joins("inner join product_version on product.id = product_version.p_id").
		Where("product.id = ?", id).
		Scan(&products)

	return products
}

func (repo userRepository) GetProductsWithCategoryId(id int) ([]dto.Product, error) {
	var products []dto.Product
	result := repo.db.GetDB().Table("product").
		Select("product.id as id, product.name as name, product.image as image, product.description as description, category.name as category_name, product.price as price").
		Joins("inner join category on category.id = product.category_id").
		Where("category.id = ?", id).Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (repo userRepository) GetProductWithId(id string) *dto.Product {
	var product *dto.Product
	repo.db.GetDB().Table("product").
		Select(
			"product.id as id,product.name as name, product.category_id as category_id, category.name as category_name, product.description as description, product.price as price, product.image as image, product.image1 as image1, product.image2 as image2").
		Joins("inner join category on category.id = product.category_id").
		Where("product.id = ?", id).Find(&product)
	return product
}
