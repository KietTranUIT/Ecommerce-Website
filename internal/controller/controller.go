package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
	u.router.LoadHTMLGlob("view/*.html")

	u.router.Static("/view/", "view")

	u.router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	u.router.GET("/account/check", CheckAccount(u))

	signup_group := u.router.Group("/signup")
	{
		signup_group.GET("/", GETSignupPage(u))
		signup_group.POST("/", SignUp(u))
	}

	verify_group := u.router.Group("/verify")
	{
		verify_group.GET("/", GETVerifyPage(u))
		verify_group.POST("/", SendVerificationCode(u))
	}

	u.router.POST("/auth", AuthenticateCode(u))
	u.router.GET("/login", GetLoginPage(u))
	u.router.POST("/login", Login(u))
}

func GetLoginPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "signin.html", nil)
	}
}

func CheckAccount(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.GetQuery("email")

		response := control.service.CheckAccount(email)

		c.Writer.Header().Set("Content-Type", "application/json")

		c.AbortWithStatusJSON(200, response)

	}
}

func GETVerifyPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "verify_signup.html", nil)
	}
}

func GETSignupPage(control UserController) gin.HandlerFunc {
	log.Println("Loi da xay ra")
	return func(c *gin.Context) {
		c.HTML(200, "signup.html", nil)
	}
}

func Login(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var data request.LoginRequest
		json.Unmarshal(reqBody, &data)

		res := control.service.Login(data)
		if !res.Status {
			c.AbortWithStatusJSON(422, res)
			return
		}
		c.AbortWithStatusJSON(200, res)
	}
}

func SendVerificationCode(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		email, _ := c.GetQuery("email")
		res := control.service.SendVerificationCode(email)

		c.Writer.Header().Set("Content-Type", "application/json")

		if res.Status == false {
			c.AbortWithStatusJSON(500, res)
			return
		}
		c.AbortWithStatusJSON(200, res)
		c.String(200, "Hello World!")
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
