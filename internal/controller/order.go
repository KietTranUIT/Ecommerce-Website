package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"user-service/internal/core/model/request"

	"github.com/gin-gonic/gin"
)

func HandleOrder(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")

		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.CreateOrderRequest
		json.Unmarshal(reqBody, &req)

		log.Println(req)
		response := control.service.CreateOrder(req)

		if !response.Status {
			c.AbortWithStatusJSON(501, response)
			return
		}
		c.AbortWithStatusJSON(200, response)
	}
}
