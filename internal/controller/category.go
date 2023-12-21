package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"user-service/internal/core/dto"
	"user-service/internal/core/model/request"

	"github.com/gin-gonic/gin"
)

// Get all category
func GetCategories(service UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := service.service.GetCategories()

		if err != nil {
			c.AbortWithStatusJSON(500, nil)
		}

		log.Println(categories)
		c.HTML(200, "CategoriesAdminPage.html", categories)
	}
}

// Create a category
func CreateCategory(service UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.CreateCategoryRequest
		json.Unmarshal(reqBody, &req)

		result := service.service.CreateCategory(req)

		if !result.Status {
			c.AbortWithStatusJSON(500, result)
		}
		c.AbortWithStatusJSON(200, result)
	}
}

// Get a category with id
func GetCategory(service UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		category, err := service.service.GetCategory(id)

		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
		}
		c.HTML(200, "UpdateCategoryAdminPage.html", category)
	}
}

// Lay ve trang them category
func GetNewCategory(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "AddCategoryAdminPage.html", nil)
	}
}

// Cap nhat thay doi category
func UpdateCategory(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var category *dto.Category
		json.Unmarshal(reqBody, &category)
		log.Println(category)

		res := control.service.UpdateCategory(category)

		if !res.Status {
			c.AbortWithStatusJSON(500, res)
		}
		c.AbortWithStatusJSON(200, res)
	}
}
