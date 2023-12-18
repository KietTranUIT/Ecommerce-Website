package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"user-service/internal/core/dto"
	"user-service/internal/core/model/request"

	"path/filepath"

	"strconv"

	"github.com/gin-gonic/gin"
)

/* REST API for admin
---------------------------------------------------
*/

//xu li yeu cau truy cap trang admin
func HandleAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("oki ban nho")
		c.String(200, "OK")

		c.HTML(http.StatusOK, "categories_admin.html", nil)
	}
}

// Tra ve trang dang nhap cho admin
func GetLoginAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "login_admin.html", nil)
	}
}

// Xu li dang nhap cho admin
func HandleLoginAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var data request.LoginRequest
		json.Unmarshal(reqBody, &data)

		res := control.service.LoginAdmin(data)

		if !res.Status {
			c.AbortWithStatusJSON(401, res)
			return
		}

		token, _ := CreateToken(data.Email)
		c.SetCookie("bear", token, 3600, "/", "localhost", false, true)
	}
}

// Lay trang hien thi danh sach san pham
func GetCategoryAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := control.service.GetCategories()
		log.Println(categories)
		c.HTML(http.StatusOK, "categories_admin.html", categories)
	}
}

//
func GetAddCategoryAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_categories_admin.html", nil)
	}
}

// Them categories
func InsertCategoryAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// lay du lieu json
		jsonData := c.PostForm("jsonData")
		var category dto.ProductCategory
		json.Unmarshal([]byte(jsonData), &category)

		category.Id = strconv.Itoa(control.service.GetLastIDCategories() + 1)
		name_image := category.Id + "_" + category.Name + "." + strings.Split(file.Filename, ".")[1]

		// luu file vao folder chi dinh
		uploadFolder := "/view/assets/image/categories"
		filePath := filepath.Join(uploadFolder, name_image)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		category.Image = filePath

		result := control.service.CreateCategory(category)
		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Them du lieu thanh cong",
		})
	}
}

func DeleteCategoryAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)

		result := control.service.DeleteCategory(id)

		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Xoa du lieu thanh cong",
		})
	}
}

func GetUpdateCategoryAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		category := control.service.GetCategoryWithId(id)

		c.HTML(http.StatusOK, "update_categories_admin.html", category)
	}
}

func UpdateCategoryAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// lay du lieu json
		jsonData := c.PostForm("jsonData")
		var category dto.ProductCategory
		json.Unmarshal([]byte(jsonData), &category)

		if file != nil {
			name_image := category.Id + "_" + category.Name + "." + strings.Split(file.Filename, ".")[1]

			// luu file vao folder chi dinh
			uploadFolder := "view/assets/image/categories"
			filePath := filepath.Join(uploadFolder, name_image)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			category.Image = "/" + filePath
		}
		log.Println(category)

		result := control.service.UpdateCategory(&category)
		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Cap nhat du lieu thanh cong",
		})
	}
}

// ----------------- Products----------------------------------
func GetProductAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		products := control.service.GetProductsForAdmin()
		log.Println(products)

		c.HTML(200, "products_admin.html", products)
	}
}

func GetAddProductAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := control.service.GetCategories()

		c.HTML(200, "add_products_admin.html", categories)
	}
}

func CreateProduct(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var data dto.Product
		json.Unmarshal(reqBody, &data)

		result := control.service.CreateProduct(&data)

		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Them du lieu thanh cong",
		})
	}
}

func GetUpdateProductAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		product := control.service.GetProductWithId(id)

		categories := control.service.GetCategories()

		c.HTML(200, "update_products_admin.html", struct {
			Product  *dto.Product
			Category []dto.ProductCategory
		}{product, categories})
	}
}

func UpdateProductAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var data dto.Product
		json.Unmarshal(reqBody, &data)

		result := control.service.UpdateProduct(&data)

		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Cap nhat du lieu thanh cong",
		})
	}
}

func GetProductDetailAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")

		products := control.service.GetProductWithId(str)
		products_version := control.service.GetProductVersion(str)

		template := struct {
			Product         *dto.Product
			Product_version []dto.ProductVersion
		}{products, products_version}
		c.HTML(200, "detail_products_admin.html", template)
	}
}

func CreateProductVersionAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// lay du lieu json
		jsonData := c.PostForm("jsonData")

		var product dto.ProductVersion
		json.Unmarshal([]byte(jsonData), &product)

		product.Id = control.service.GetLastIdProductVersion() + 1
		log.Println(product)
		name_image := strconv.Itoa(product.Id) + "." + strings.Split(file.Filename, ".")[1]

		// luu file vao folder chi dinh
		uploadFolder := "view/assets/image/products"
		filePath := filepath.Join(uploadFolder, name_image)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		product.Image = "/" + filePath

		result := control.service.CreateProductVersion(&product)
		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}

		// var inventory *dto.ProductInventory
		// log.Println(product.Id)
		// inventory.Product_id = product.Id
		// quantity, _ := strconv.Atoi(product.Inventory)
		// inventory.Quantity = quantity
		// log.Println(inventory)
		var p_id = product.Id
		quantity, _ := strconv.Atoi(product.Inventory)

		inventory := dto.ProductInventory{
			Product_id: p_id,
			Quantity:   quantity,
		}

		result_new := control.service.CreateProductInventory(&inventory)
		if !result_new {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Them du lieu thanh cong",
		})
	}
}

func GetAddProductVersionAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		c.HTML(200, "add_productversion_admin.html", gin.H{"Id": id})
	}
}

func GetUpdateProductVersionAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)
		product := control.service.GetProductVersionWithId(id)

		c.HTML(200, "update_productversion_admin.html", product)
	}
}

func UpdateProductVersionAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// lay du lieu json
		jsonData := c.PostForm("jsonData")
		var product dto.ProductVersion
		json.Unmarshal([]byte(jsonData), &product)

		if file != nil {
			name_image := strconv.Itoa(product.Id) + "." + strings.Split(file.Filename, ".")[1]

			// luu file vao folder chi dinh
			uploadFolder := "view/assets/image/products"
			filePath := filepath.Join(uploadFolder, name_image)
			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product.Image = "/" + filePath
		}

		result := control.service.UpdateProductVersion(&product)
		if !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}

		quantity, _ := strconv.Atoi(product.Inventory)
		result1 := control.service.UpdateProductInventory(product.Id, quantity)
		if !result1 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Cap nhat du lieu thanh cong",
		})
	}
}

func DeleteProductVersion(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)

		if result := control.service.DeleteProductVersion(id); !result {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Xoa du lieu thanh cong",
		})
	}
}

func GetOrderAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		orders := control.service.GetOrderAdminPage()

		c.HTML(200, "orders_admin.html", orders)
	}
}

func GetOrderDetailAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)

		order := control.service.GetOrderWithId(id)
		orders_detail := control.service.GetOrderDetail(id)

		c.HTML(200, "detail_orders_admin.html", struct {
			Order       *dto.Order
			OrderDetail []dto.OrderDetail
		}{order, orders_detail})
	}
}
