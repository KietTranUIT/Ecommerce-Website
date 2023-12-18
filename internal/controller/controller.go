package controller

import (
	"user-service/internal/core/port/service"
	"user-service/internal/middleware"

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

	// For Admin API
	admin_group := u.router.Group("/admin", middleware.AuthenticateAdmin())
	{
		admin_group.GET("/home", HandleAdmin(u))
		admin_group.GET("/categories", GetCategoryAdminPage(u))
		admin_group.GET("/categories/new", GetAddCategoryAdminPage(u))
		admin_group.POST("/categories", InsertCategoryAdmin(u))
		admin_group.DELETE("/categories/:id", DeleteCategoryAdmin(u))
		admin_group.GET("/categories/update/:id", GetUpdateCategoryAdminPage(u))
		admin_group.PUT("/categories/update/:id", UpdateCategoryAdmin(u))

		admin_group.GET("/products", GetProductAdminPage(u))
		admin_group.GET("/products/new", GetAddProductAdminPage(u))
		admin_group.POST("/products", CreateProduct(u))
		admin_group.GET("/products/update/:id", GetUpdateProductAdminPage(u))
		admin_group.PUT("/products/update/:id", UpdateProductAdmin(u))
		admin_group.GET("/products/:id", GetProductDetailAdminPage(u))
		admin_group.GET("/products/:id/new", GetAddProductVersionAdminPage(u))
		admin_group.POST("/products/:id/products_version", CreateProductVersionAdmin(u))
		admin_group.GET("/products_version/update/:id", GetUpdateProductVersionAdminPage(u))
		admin_group.PUT("/products_version/update/:id", UpdateProductVersionAdmin(u))
		admin_group.DELETE("/products_version/:id", DeleteProductVersion(u))
		admin_group.GET("/orders", GetOrderAdminPage(u))
		admin_group.GET("/orders/:id", GetOrderDetailAdminPage(u))
	}

	admin_login := u.router.Group("/admin/login")
	{
		admin_login.GET("/", GetLoginAdminPage(u))
		admin_login.POST("/", HandleLoginAdmin(u))
	}

	user_group := u.router.Group("/account")
	{
		user_group.POST("/addresses", HandleCreateUserAddress(u))
	}
}
