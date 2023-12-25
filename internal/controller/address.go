package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"

	"user-service/internal/core/model/request"

	"github.com/gin-gonic/gin"
)

// Ham goi service tao mot dia chi moi cho user
func HandleCreateUserAddress(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		user_email, _ := c.Get("userId")

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.CreateUserAddressRequest
		json.Unmarshal(reqBody, &req)
		req.User_email = user_email.(string)

		response := control.service.CreateUserAddress(req)

		if !response.Status {
			c.AbortWithStatusJSON(501, response)
			return
		}
		c.AbortWithStatusJSON(200, response)
	}
}

// Ham goi service xoa mot dia chi cua user
func HandleDeleteUserAddress(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		id, _ := strconv.Atoi(c.Param("id"))

		req := request.DeleteUserAddressRequest{Id: id}
		response := control.service.DeleteUserAddress(req)
		if !response.Status {
			c.AbortWithStatusJSON(501, response)
			return
		}
		c.AbortWithStatusJSON(200, response)
	}
}

// Ham goi service cap nhat thong tin dia chi cua user
func HandleUpdateUserAddress(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		id, _ := strconv.Atoi(c.Param("id"))

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.EditUserAddressRequest
		json.Unmarshal(reqBody, &req)
		req.Id = id

		response := control.service.EditUserAddress(req)

		if !response.Status {
			c.AbortWithStatusJSON(501, response)
			return
		}
		c.AbortWithStatusJSON(200, response)
	}
}

// Ham goi service lay ve cac dia chi cua mot user cu the
func GetUserAddress(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_email, _ := c.Get("userId")

		addresses, err := control.service.GetUserAddress(user_email.(string))
		log.Println("OK")

		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
		}
		c.HTML(200, "Address.html", addresses)
	}
}
