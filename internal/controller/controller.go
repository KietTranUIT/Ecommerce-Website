package controller

import (
	"encoding/json"
	"io/ioutil"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
	"user-service/internal/core/port/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router  *gin.Engine
	service service.UserService
}

func NewUserController(router *gin.Engine, service service.UserService) UserController {
	return UserController{
		router:  router,
		service: service,
	}
}

func (u UserController) InitRouter() {
	u.router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	u.router.POST("/signup", SignUp(u))
	u.router.POST("/verify", SendVerificationCode(u))
	u.router.POST("/auth", AuthenticateCode(u))
}

func SendVerificationCode(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.GetQuery("email")
		res := control.service.SendVerificationCode(email)

		if res.Status == false {
			c.AbortWithStatusJSON(500, res)
			return
		}
		c.AbortWithStatusJSON(200, res)
	}
}

func SignUp(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read data from json data in body request
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var data request.SignUpRequest
		json.Unmarshal(reqBody, &data)

		// call Sign up service and receive response return
		var res *response.Response
		res = control.service.SignUp(data)

		if res.Status == false {
			if res.Error_code == error_code.Duplicate_code {
				c.AbortWithStatusJSON(422, res)
				return
			}
			c.AbortWithStatusJSON(500, res)
			return
		}
		c.AbortWithStatusJSON(200, res)

	}
}

func AuthenticateCode(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var data request.AuthenticateRequest
		json.Unmarshal(reqBody, &data)

		res := control.service.AuthenticateCode(data)

		if res.Status == false {
			c.AbortWithStatusJSON(422, res)
			return
		}
		c.AbortWithStatusJSON(200, res)
	}
}
