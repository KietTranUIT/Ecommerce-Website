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

	// Load nhung file html, css, js
	u.router.LoadHTMLGlob("view/*.html")
	u.router.Static("/view/", "view")

	// Su dung middleware hien thi ra man hinh request http
	u.router.Use(gin.Logger())

	/*--------------User routes------------------*/
	// route truy cap trang chu
	u.router.GET("/", HomePage(u))

	// route kiem tra account co ton tai hay khong
	u.router.GET("/account/check", CheckAccount(u))

	// route dung dang ki
	signup_group := u.router.Group("/signup")
	{
		signup_group.GET("/", GETSignupPage(u))
		signup_group.POST("/", SignUp(u))
	}

	// route dung xac thuc code gui ve email
	verify_group := u.router.Group("/verify")
	{
		verify_group.GET("/", GETVerifyPage(u))
		verify_group.POST("/", SendVerificationCode(u))
	}
	u.router.POST("/auth", AuthenticateCode(u))

	// route dung de dang nhap
	u.router.GET("/login", GetLoginPage(u))
	u.router.POST("/login", Login(u))

	u.router.POST("/orders", HandleOrder(u))
	u.router.GET("/home", HandleHomePage(u))
	u.router.GET("/products/:id", GetProductDetail(u))

	u.router.GET("/products", GetProducts(u))

	u.router.GET("/categories/:id", GetProductsOfCategory(u))

	u.router.GET("/cart", GetCart(u))

	u.router.GET("/checkout", middleware.AuthenticateUser(), GetCheckout(u))
	u.router.POST("/checkout", middleware.AuthenticateUser(), HandleCheckout(u))

	users_group := u.router.Group("/user", middleware.AuthenticateUser())
	{
		users_group.GET("/address", GetUserAddress(u))
		users_group.POST("/address", HandleCreateUserAddress(u))
		users_group.DELETE("/address/:id", HandleDeleteUserAddress(u))
		users_group.PUT("/address/:id", HandleUpdateUserAddress(u))
		users_group.GET("/", GetProfileUser(u))
	}

	u.router.DELETE("/logout", Logout(u))

	/*--------------Admin routes------------------*/
	admin_group := u.router.Group("/admin", middleware.AuthenticateAdmin())
	{
		// route truy cap trang chu Admin
		admin_group.GET("/", HandleAdmin(u))

		// route tra ve doanh so luong don hang
		admin_group.GET("/sales/filter", GetTotalSales(u))

		// // rout
		// admin_group.GET("/ordersrecently", GetOrdersRecently(u))
		// admin_group.GET("/topproducts", GetTopProducts(u))

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

}
