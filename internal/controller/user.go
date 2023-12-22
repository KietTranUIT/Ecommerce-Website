package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"

	"github.com/gin-gonic/gin"
)

func HomePage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "HomePage_v1.html", nil)
	}
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

func HandleHomePage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		products := control.service.GetProductForHomePage()

		c.AbortWithStatusJSON(200, products)
	}
}

func HandleCheckout(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		//user_email, err := c.Get("user_email")

		// if !err {
		// 	c.AbortWithStatus(401)
		// }
		log.Println("OK")
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.CreateOrderRequest
		json.Unmarshal(reqBody, &req)
		req.User_email = "kiettranuit@gmail.com"

		response := control.service.CreateOrder(req)

		if !response.Status {
			c.AbortWithStatusJSON(500, response)
		}
		c.AbortWithStatusJSON(200, response)

	}
}
