package service

import (
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
)

// Logic service get all category
func (service userService) GetCategories() ([]dto.Category, error) {
	categories, err := service.repo.GetAllCategory()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (service userService) CreateCategory(req request.CreateCategoryRequest) *response.Response {
	category := &dto.Category{
		Name:        req.Name,
		Person:      req.Person,
		Description: req.Description,
	}
	err := service.repo.CreateCategory(category)

	if err != nil {
		return CreateFailResponse(error_code.CreateCategoryFail, err.Error())
	}
	return CreateSuccessResponse(error_code.Success, "")
}

func (service userService) GetCategory(id int) (*dto.Category, error) {
	category, err := service.repo.GetCategory(id)

	if err != nil {
		return nil, err
	}
	return category, nil
}

func (service userService) UpdateCategory(category *dto.Category) *response.Response {
	result := service.repo.UpdateCategory(category)
	if result != nil {
		return CreateFailResponse(error_code.UpdateCategoryFail, result.Error())
	}
	return CreateSuccessResponse(error_code.Success, "")
}
