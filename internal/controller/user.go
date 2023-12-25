package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"user-service/internal/core/common/util"
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"

	"github.com/gin-gonic/gin"
)

func HomePage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "HomePage.html", nil)
	}
}

func GetLoginPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "Login.html", nil)
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
		c.HTML(200, "VerifySignUp.html", nil)
	}
}

func GETSignupPage(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "SignUp.html", nil)
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
		token, _ := util.CreateToken(data.Email)
		c.SetCookie("user-token", token, 3600, "/", "https://377c-42-115-60-125.ngrok-free.app", false, true)
		log.Println(res)
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
		user_email, _ := c.Get("userId")

		// if !err {
		// 	c.AbortWithStatus(401)
		// }
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		var req request.CreateOrderRequest
		json.Unmarshal(reqBody, &req)
		req.User_email = user_email.(string)

		response := control.service.CreateOrder(req)

		if !response.Status {
			c.AbortWithStatusJSON(500, response)
		}
		c.AbortWithStatusJSON(200, response)

	}
}

func GetCart(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(200, "Cart.html", nil)
	}
}

func GetCheckout(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_email, _ := c.Get("userId")

		address, _ := control.service.GetUserAddress(user_email.(string))

		payment_method := control.service.GetPaymentMethod()

		c.HTML(200, "Checkout.html", struct {
			Address       []dto.UserAddress
			PaymentMethod []dto.PaymentMethod
		}{address, payment_method})
	}
}

func GetProfileUser(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		user_email, err := c.Get("userId")

		if !err {
			return
		}
		orders := control.service.GetOrderWithEmail(user_email.(string))

		c.HTML(200, "Profile.html", orders)
	}
}

func Logout(control UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("user-token", "", -1, "/", "https://377c-42-115-60-125.ngrok-free.app", false, true)
		c.String(200, "")
	}
}
