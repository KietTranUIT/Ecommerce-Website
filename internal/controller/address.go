package controller

import (
	"encoding/json"
	"io/ioutil"

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
