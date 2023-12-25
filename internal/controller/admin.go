package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"user-service/internal/core/dto"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"

	"strconv"

	"user-service/internal/core/common/util"

	"github.com/gin-gonic/gin"
)

//xu li yeu cau truy cap trang admin
func HandleAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		summary, products, orders := control.service.HandleAdmin()

		c.HTML(http.StatusOK, "AdminPage.html", struct {
			Summary  *dto.DataSales
			Products []dto.Product
			Orders   []dto.Order
		}{summary, products, orders})
	}
}

// Tra ve trang dang nhap cho admin
func GetLoginAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "LoginAdminPage.html", nil)
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

		token, _ := util.CreateToken(data.Email)
		c.SetCookie("admin-token", token, 3600, "/admin", "localhost", false, true)
		c.AbortWithStatusJSON(200, res)
	}
}

// Lay trang hien thi danh sach san pham
// func GetCategoryAdminPage(control UserController) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		categories := control.service.GetCategories()
// 		c.HTML(http.StatusOK, "CategoriesAdminPage.html", categories)
// 	}
// }

//

// Them categories
// func InsertCategoryAdmin(control UserController) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		file, err := c.FormFile("file")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// lay du lieu json
// 		jsonData := c.PostForm("jsonData")
// 		var category dto.ProductCategory
// 		json.Unmarshal([]byte(jsonData), &category)

// 		category.Id = strconv.Itoa(control.service.GetLastIDCategories() + 1)
// 		name_image := category.Id + "_" + category.Name + "." + strings.Split(file.Filename, ".")[1]

// 		// luu file vao folder chi dinh
// 		uploadFolder := "view/assets/image/categories"
// 		filePath := filepath.Join(uploadFolder, name_image)
// 		if err := c.SaveUploadedFile(file, filePath); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		category.Image = "/" + filePath

// 		result := control.service.CreateCategory(category)
// 		if !result {
// 			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Them du lieu thanh cong",
// 		})
// 	}
// }

func DeleteCategoryAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)

		result := control.service.DeleteCategory(id)
		log.Println(result)

		if !result.Status {
			c.AbortWithStatusJSON(500, result)
			return
		}
		c.AbortWithStatusJSON(500, result)
	}
}

func GetUpdateCategoryAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		category := control.service.GetCategoryWithId(id)

		c.HTML(http.StatusOK, "UpdateCategoryAdminPage.html", category)
	}
}

// func UpdateCategoryAdmin(control UserController) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		file, _ := c.FormFile("file")
// 		// if err != nil {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// 	return
// 		// }

// 		// lay du lieu json
// 		jsonData := c.PostForm("jsonData")
// 		var category dto.ProductCategory
// 		json.Unmarshal([]byte(jsonData), &category)

// 		if file != nil {
// 			name_image := category.Id + "_" + category.Name + "." + strings.Split(file.Filename, ".")[1]

// 			// luu file vao folder chi dinh
// 			uploadFolder := "view/assets/image/categories"
// 			filePath := filepath.Join(uploadFolder, name_image)
// 			if err := c.SaveUploadedFile(file, filePath); err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 				return
// 			}
// 			category.Image = "/" + filePath
// 		}

// 		result := control.service.UpdateCategory(&category)
// 		if !result {
// 			c.JSON(http.StatusInternalServerError, gin.H{"message": "Loi he thong"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Cap nhat du lieu thanh cong",
// 		})
// 	}
// }

// ----------------- Products----------------------------------

func DeleteProduct(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)

		result := control.service.DeleteProduct(id)

		if !result.Status {
			c.AbortWithStatusJSON(500, result)
			return
		}
		c.AbortWithStatusJSON(200, result)
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
		c.HTML(200, "DetailProductAdminPage.html", template)
	}
}

func CreateProductVersionAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var product dto.ProductVersion
		json.Unmarshal(reqBody, &product)

		product.Id = control.service.GetLastIdProductVersion() + 1

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
		log.Println("OK")

		c.HTML(200, "AddProductVersionAdminPage.html", gin.H{"Id": id})
	}
}

func GetUpdateProductVersionAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)
		product := control.service.GetProductVersionWithId(id)

		c.HTML(200, "UpdateProductVersionAdminPage.html", product)
	}
}

func UpdateProductVersionAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var product dto.ProductVersion
		json.Unmarshal(reqBody, &product)
		log.Println(product)

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
		c.HTML(200, "OrdersAdminPage.html", orders)
	}
}

func GetOrderDetailAdminPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		str := c.Param("id")
		id, _ := strconv.Atoi(str)

		order := control.service.GetOrderWithId(id)
		orders_detail := control.service.GetOrderDetail(id)

		c.HTML(200, "OrderDetailAdminPage.html", struct {
			Order       *dto.Order
			OrderDetail []dto.OrderDetail
		}{order, orders_detail})
	}
}

func HandleLogoutAdmin(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("admin-token", "", -1, "/admin", "localhost", false, true)
		c.String(200, "")
	}
}

func GetTotalSales(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		type_query := c.Query("type")

		var response *response.Response

		if type_query == "day" {
			response = control.service.GetTotalSalesDayNow()
		} else if type_query == "week" {
			response = control.service.GetTotalSalesWeekNow()
		} else {
			response = control.service.GetTotalSalesMonthNow()
		}

		if !response.Status {
			c.AbortWithStatusJSON(500, response)
		}
		c.AbortWithStatusJSON(200, response)
	}
}

func GetOrdersRecently(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := control.service.GetOrdersRecently()

		if !response.Status {
			c.AbortWithStatusJSON(500, response)
		}
		c.AbortWithStatusJSON(200, response)
	}
}

func GetTopProducts(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := control.service.GetTopProducts()

		if !response.Status {
			c.AbortWithStatusJSON(500, response)
		}
		c.AbortWithStatusJSON(200, response)
	}
}
