package controller

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"user-service/internal/core/dto"

	"github.com/gin-gonic/gin"
)

func GetProductDetail(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str_id := c.Param("id")
		id, _ := strconv.Atoi(str_id)

		products := control.service.GetProductData(id)

		c.AbortWithStatusJSON(200, products)
	}
}

// Get products
func GetProducts(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := control.service.GetProductForHomePage()

		if !response.Status {
			c.AbortWithStatusJSON(500, response)
			return
		}
		c.AbortWithStatusJSON(200, response)
	}
}

// Get products with category id
func GetProductsCategory(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("category_id"))

		res := control.service.GetProductsWithCategoryId(id)
		if !res.Status {
			c.AbortWithStatusJSON(500, res)
		}
		c.AbortWithStatusJSON(200, res)
	}
}

// Get products admin
func GetProductAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, _ := control.service.GetCategories()
		products := control.service.GetProductsForAdmin()

		c.HTML(200, "ProductsAdminPage.html", struct {
			Product  []dto.Product
			Category []dto.Category
		}{products, categories})
	}
}

func GetNewProduct(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := control.service.GetCategories()

		if err != nil {
			c.AbortWithStatusJSON(500, err)
		}

		c.HTML(200, "AddProductAdminPage.html", categories)
	}
}

func CreateProduct(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		file1, err1 := c.FormFile("file1")
		file2, err2 := c.FormFile("file2")

		if err != nil || err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// lay du lieu json
		jsonData := c.PostForm("jsonData")
		var product dto.Product
		json.Unmarshal([]byte(jsonData), &product)

		product.Id = control.service.GetLastIdProduct() + 1
		name_image := strconv.Itoa(product.Id) + "_image" + "." + strings.Split(file.Filename, ".")[1]
		name_image1 := strconv.Itoa(product.Id) + "_image1" + "." + strings.Split(file1.Filename, ".")[1]
		name_image2 := strconv.Itoa(product.Id) + "_image2" + "." + strings.Split(file2.Filename, ".")[1]

		var name_folder = strconv.Itoa(product.Id)
		// luu file vao folder chi dinh
		uploadFolder := "view/assets/image/products/" + name_folder
		filePath := filepath.Join(uploadFolder, name_image)
		filePath1 := filepath.Join(uploadFolder, name_image1)
		filePath2 := filepath.Join(uploadFolder, name_image2)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := c.SaveUploadedFile(file1, filePath1); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := c.SaveUploadedFile(file2, filePath2); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		product.Image = "/" + filePath
		product.Image1 = "/" + filePath1
		product.Image2 = "/" + filePath2

		result := control.service.CreateProduct(&product)
		if !result {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "Them du lieu thanh cong",
		})
	}
}

func GetUpdateProductAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		product := control.service.GetProductWithId(id)

		categories, err := control.service.GetCategories()
		if err != nil {
			c.AbortWithStatusJSON(500, err)
		}

		c.HTML(200, "UpdateProductAdminPage.html", struct {
			Product  *dto.Product
			Category []dto.Category
		}{product, categories})
	}
}

func UpdateProductAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, _ := c.FormFile("file")
		file1, _ := c.FormFile("file1")
		file2, _ := c.FormFile("file2")

		// lay du lieu json
		jsonData := c.PostForm("jsonData")
		var product dto.Product
		json.Unmarshal([]byte(jsonData), &product)
		var name_folder = strconv.Itoa(product.Id)

		if file != nil {
			name_image := strconv.Itoa(product.Id) + "_image" + "." + strings.Split(file.Filename, ".")[1]

			// luu file vao folder chi dinh
			uploadFolder := "view/assets/image/products/" + name_folder
			filePath := filepath.Join(uploadFolder, name_image)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product.Image = "/" + filePath
		}
		if file1 != nil {
			name_image := strconv.Itoa(product.Id) + "_image1" + "." + strings.Split(file.Filename, ".")[1]

			// luu file vao folder chi dinh
			uploadFolder := "view/assets/image/products/" + name_folder
			filePath := filepath.Join(uploadFolder, name_image)
			if err := c.SaveUploadedFile(file1, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product.Image1 = "/" + filePath
		}
		if file2 != nil {
			name_image := strconv.Itoa(product.Id) + "_image2" + "." + strings.Split(file.Filename, ".")[1]

			// luu file vao folder chi dinh
			uploadFolder := "view/assets/image/products/" + name_folder
			filePath := filepath.Join(uploadFolder, name_image)
			if err := c.SaveUploadedFile(file2, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product.Image2 = "/" + filePath
		}

		result := control.service.UpdateProduct(&product)
		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Cap nhat du lieu thanh cong",
		})
	}
}

func GetProductsOfCategory(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		products := control.service.GetProductsWithCategoryIdV1(id)

		c.HTML(200, "ProductsCategory.html", products)
	}
}
