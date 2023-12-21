package repository

import (
	"user-service/internal/core/dto"
)

// Get category records from category table
func (repo userRepository) GetAllCategory() ([]dto.Category, error) {
	var categories []dto.Category
	result := repo.db.GetDB().Table("category").Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

// Insert a category to database
func (repo userRepository) CreateCategory(category *dto.Category) error {
	result := repo.db.GetDB().Table("category").Select("name", "description", "person").Create(&category)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get category with id
func (repo userRepository) GetCategory(id int) (*dto.Category, error) {
	var category *dto.Category
	result := repo.db.GetDB().Table("category").Where("id = ?", id).Find(&category)

	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (repo userRepository) UpdateCategory(category *dto.Category) error {
	result := repo.db.GetDB().Table("category").Where("id = ?", category.Id).Updates(
		map[string]interface{}{
			"name":        category.Name,
			"description": category.Description,
			"person":      category.Person,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
