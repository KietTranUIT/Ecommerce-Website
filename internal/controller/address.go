package controller

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"user-service/internal/core/model/request"

	"github.com/gin-gonic/gin"
)

func HandleCreateUserAddress(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.CreateUserAddressRequest
		json.Unmarshal(reqBody, &req)

		response := control.service.CreateUserAddress(req)

		if !response.Status {
			c.AbortWithStatusJSON(501, response)
			return
		}
		c.AbortWithStatusJSON(200, response)
	}
}

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

func HandleUpdateUserAddress(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.EditUserAddressRequest
		json.Unmarshal(reqBody, &req)

		response := control.service.EditUserAddress(req)

		if !response.Status {
			c.AbortWithStatusJSON(501, response)
			return
		}
		c.AbortWithStatusJSON(200, response)
	}
}
