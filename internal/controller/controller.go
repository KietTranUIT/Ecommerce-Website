package controller

import (
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
	u.router.LoadHTMLGlob("view/admin/*.html")

	u.router.Static("/view/", "view")

	u.router.Use(gin.Logger())

	u.router.GET("/", HomePage(u))

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
	admin_group := u.router.Group("/admin")
	{
		admin_group.GET("/", HandleAdmin(u))
		admin_group.GET("/sales/filter", GetTotalSales(u))
		admin_group.GET("/ordersrecently", GetOrdersRecently(u))
		admin_group.GET("/topproducts", GetTopProducts(u))

		// route category
		admin_group.GET("/categories", GetCategories(u))
		admin_group.POST("/categories", CreateCategory(u))
		admin_group.GET("/categories/update/:id", GetCategory(u))
		admin_group.GET("/categories/new", GetNewCategory(u))
		admin_group.PUT("/categories/:id", UpdateCategory(u))
		admin_group.DELETE("/categories/:id", DeleteCategoryAdmin(u))
		admin_group.GET("/categories/:id/products", GetProductsCategory(u))
		admin_group.GET("/products/new", GetNewProduct(u))
		admin_group.GET("/products/update/:id", GetUpdateProductAdminPage(u))
		admin_group.PUT("/products/update/:id", UpdateProductAdmin(u))

		//admin_group.GET("/categories/update/:id", GetUpdateCategoryAdminPage(u))
		//admin_group.PUT("/categories/update/:id", UpdateCategoryAdmin(u))

		// route product
		admin_group.GET("/products", GetProductAdminPage(u))
		admin_group.GET("/products/filter", GetProductsCategory(u))
		admin_group.POST("/products", CreateProduct(u))

		admin_group.DELETE("/products/:id", DeleteProduct(u))
		admin_group.GET("/products/:id", GetProductDetailAdminPage(u))
		admin_group.GET("/products/:id/new", GetAddProductVersionAdminPage(u))
		admin_group.POST("/products/:id/products_version", CreateProductVersionAdmin(u))
		admin_group.GET("/products_version/update/:id", GetUpdateProductVersionAdminPage(u))
		admin_group.PUT("/products_version/update/:id", UpdateProductVersionAdmin(u))
		admin_group.DELETE("/products_version/:id", DeleteProductVersion(u))
		admin_group.GET("/orders", GetOrderAdminPage(u))
		admin_group.GET("/orders/:id", GetOrderDetailAdminPage(u))
		admin_group.DELETE("/logout", HandleLogoutAdmin(u))
	}
	u.router.GET("/admin/login", GetLoginAdminPage(u))
	u.router.POST("/admin/login", HandleLoginAdmin(u))

	user_group := u.router.Group("/account")
	{
		user_group.POST("/addresses", HandleCreateUserAddress(u))
		user_group.DELETE("/addresses/:id", HandleDeleteUserAddress(u))
		user_group.PUT("/addresses/:id", HandleUpdateUserAddress(u))
	}

	u.router.POST("/orders", HandleOrder(u))
	u.router.GET("/home", HandleHomePage(u))
	u.router.GET("/products/:id", GetProductDetail(u))

	u.router.GET("/products", GetProducts(u))
	u.router.POST("/checkout", HandleCheckout(u))

	u.router.GET("/categories/:id", GetProductsOfCategory(u))

	u.router.GET("/cart", GetCart(u))
}
