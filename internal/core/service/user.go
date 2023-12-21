package service

import (
	"log"
	"strings"
	"time"
	"user-service/internal/core/common/util"
	"user-service/internal/core/dto"
	"user-service/internal/core/entity/error_code"
	"user-service/internal/core/entity/mail"
	"user-service/internal/core/model/request"
	"user-service/internal/core/model/response"
	"user-service/internal/core/port/repository"
	"user-service/internal/core/port/service"
)

const (
	duplicateEntry = "Duplicate entry"
)

type userService struct {
	repo repository.UserRepository
	mail.MailService
}

func NewUserService(repo repository.UserRepository) service.UserService {
	mail, err := mail.NewMailService("/home/kiettran/IT/ecommerce-project/internal/core/entity/mail")
	if err != nil {
		log.Println("error config mail service ", err.Error())
	}

	return userService{
		repo:        repo,
		MailService: mail,
	}
}

func (service userService) CheckAccount(email string) *response.Response {
	user := service.repo.GetUserWithEmail(email)

	if user == nil {
		return CreateFailResponse(error_code.Empty, error_code.NotExistUser_msg)
	}
	return CreateSuccessResponse(error_code.Duplicate_code, error_code.DuplicateUserEmail_msg)
}

func (service userService) Login(req request.LoginRequest) *response.Response {
	user := service.repo.GetUserWithEmail(req.Email)
	log.Println(user)

	if user == nil {
		return CreateFailResponse(error_code.LoginError, error_code.NotExistUser_msg)
	}

	if user.Password != req.Password {
		return CreateFailResponse(error_code.LoginError, error_code.WrongPassword)
	}

	return CreateSuccessResponse(error_code.LoginSuccess, error_code.LoginSuccess_msg, user)
}

// Send code to email user
func (service userService) SendVerificationCode(email string) *response.Response {
	code := util.RandomCode()

	// Select verify data from database
	verify := service.repo.GetVerificationWithEmailAndType(email, "sign up")

	if verify != nil {
		if verify.Expire_at.Before(time.Now()) {
			service.repo.UpdateVerificationEmail(
				dto.VerificationEmail{
					Email:     email,
					Code:      code,
					Type:      "sign up",
					Expire_at: util.ExpireAt(),
				},
			)
			goto label
		}
		return CreateFailResponse(error_code.Duplicate_code, error_code.DuplicateEmailVerification_msg)
	}

	if err := service.repo.CreateVerificationMail(
		&dto.VerificationEmail{
			Email:     email,
			Code:      code,
			Type:      "sign up",
			Expire_at: util.ExpireAt(),
		},
	); err != nil {
		return CreateFailResponse(error_code.InternalError_code, error_code.InternalError_msg)

	}

label:
	mail_data := mail.CreateVerificationMail([]string{email}, code)

	if err := service.SendMail(mail_data); err != nil {
		return CreateFailResponse(error_code.SendCodeFailError_code, error_code.SendCodeFailError_msg)
	}

	return CreateSuccessResponse(error_code.SendCodeSuccess, error_code.SendCodeSuccess_msg)
}

func (service userService) AuthenticateCode(req request.AuthenticateRequest) *response.Response {
	verify := service.repo.GetVerificationWithEmailAndType(req.Email, req.Kind)

	if verify.Expire_at.Before(time.Now()) {
		return CreateFailResponse(error_code.ExpiredCode, error_code.ExpiredCode_msg)
	}

	if verify.Code != req.Code {
		return CreateFailResponse(error_code.FailedAuthentication, error_code.FailedAuthentication_msg)
	}
	service.repo.UpdateStatusVerificationEmail(req.Email, req.Kind)
	return CreateSuccessResponse(error_code.SuccessAuthentication, error_code.SuccessAuthentication_msg)
}

func (service userService) SignUp(req request.SignUpRequest) *response.Response {
	response := service.AuthenticateCode(request.AuthenticateRequest{
		Email: req.Email,
		Kind:  "sign up",
		Code:  req.Code,
	})

	if !response.Status {
		return response
	}

	// verify := service.repo.GetVerificationWithEmailAndType(req.Email, "sign up")

	// if verify.Status == false {
	// 	return CreateFailResponse(error_code.NotAuthenticatedError_code, error_code.NotAuthenticated_msg)
	// }

	// Create a user object
	user := dto.UserDTO{
		Email:       req.Email,
		Password:    req.Password,
		First_name:  req.FirstName,
		Last_name:   req.LastName,
		Gender:      req.Gender,
		Created_at:  time.Now(),
		Modified_at: time.Now(),
		Active:      true,
	}

	// Check errors when insert a user into database
	if err := service.repo.CreateUser(&user); err != nil {
		if strings.Contains(err.Error(), duplicateEntry) {
			return CreateFailResponse(error_code.Duplicate_code, error_code.DuplicateUserEmail_msg)
		}
		return CreateFailResponse(error_code.InternalError_code, error_code.InternalError_msg)
	}

	//Create a user address object
	// user_address := dto.UserAddress{
	// 	User_id:   user.Id,
	// 	Telephone: req.Telephone,
	// 	Address:   req.Address,
	// }

	// if err := service.repo.CreateUserAddress(&user_address); err != nil {
	// 	if strings.Contains(err.Error(), duplicateEntry) {
	// 		return CreateFailResponse(error_code.Duplicate_code, error_code.DuplicateUserTelephone_msg)
	// 	}
	// 	return CreateFailResponse(error_code.InternalError_code, error_code.InternalError_msg)
	// }

	return CreateSuccessResponse(error_code.Signup_success, "success")
}

func CreateFailResponse(err_code error_code.Error_code, err_msg string, data ...any) *response.Response {
	return &response.Response{
		Data:       data,
		Status:     false,
		Error_code: err_code,
		Error_msg:  err_msg,
	}
}

func CreateSuccessResponse(err_code error_code.Error_code, err_msg string, data ...any) *response.Response {
	return &response.Response{
		Data:       data,
		Status:     true,
		Error_code: err_code,
		Error_msg:  err_msg,
	}
}

// Logic service for Admin ----------------------------------------------------------
func (service userService) LoginAdmin(req request.LoginRequest) *response.Response {
	admin := service.repo.GetAdmin(req.Email)

	if admin == nil {
		return CreateFailResponse(error_code.LoginError, error_code.LoginAdminFail_msg1)
	}

	if req.Password != admin.Password {
		return CreateFailResponse(error_code.LoginError, error_code.LoginAdminFail_msg2)
	}

	return CreateSuccessResponse(error_code.LoginSuccess, error_code.LoginSuccess_msg)
}

// func (service userService) GetCategories() []dto.ProductCategory {
// 	return service.repo.GetProductCategories()
// }

func (service userService) GetLastIDCategories() int {
	return service.repo.GetLastIDCategories()
}

// func (service userService) CreateCategory(category dto.ProductCategory) bool {
// 	result := service.repo.InsertCategory(category)
// 	if result != nil {
// 		return false
// 	}
// 	return true
// }

func (service userService) DeleteCategory(id int) *response.Response {
	result := service.repo.DeleteCategory(id)
	if result != nil {
		return CreateFailResponse(error_code.DeleteCategoryFail, result.Error())
	}
	return CreateSuccessResponse(error_code.Success, "")
}

func (service userService) GetCategoryWithId(id string) *dto.ProductCategory {
	return service.repo.GetCategoryWithId(id)

}

func (service userService) GetProductsForAdmin() []dto.Product {
	return service.repo.GetProductForAdmin()
}

func (service userService) GetProductWithId(id string) *dto.Product {
	return service.repo.GetProductWithId(id)
}

func (service userService) UpdateProduct(product *dto.Product) bool {
	result := service.repo.UpdateProduct(product)

	if result != nil {
		return false
	}
	return true
}

func (service userService) GetProductVersion(id string) []dto.ProductVersion {
	var product []dto.ProductVersion
	product = service.repo.GetProductVersion(id)
	return product
}

func (service userService) GetLastIdProduct() int {
	return service.repo.GetLastIdProduct()
}

func (service userService) GetLastIdProductVersion() int {
	return service.repo.GetLastIdProductVersion()
}

func (service userService) DeleteProduct(id int) *response.Response {
	if result := service.repo.DeleteAllProductVersion(id); result != nil {
		return CreateFailResponse(error_code.DeleteProductFail, result.Error())
	}

	if result := service.repo.DeleteProduct(id); result != nil {
		return CreateFailResponse(error_code.DeleteProductFail, result.Error())
	}
	return CreateSuccessResponse(error_code.Success, "")
}

func (service userService) CreateProductVersion(product *dto.ProductVersion) bool {
	result := service.repo.CreateProductVersion(product)

	if result != nil {
		return false
	}
	return true
}

func (service userService) CreateProductInventory(inventory *dto.ProductInventory) bool {
	result := service.repo.CreateProductInventory(inventory)

	if result != nil {
		return false
	}
	return true
}

func (service userService) UpdateProductInventory(id int, quantity int) bool {
	result := service.repo.UpdateProductInventory(id, quantity)

	if result != nil {
		return false
	}
	return true
}

func (service userService) UpdateProductVersion(product *dto.ProductVersion) bool {
	result := service.repo.UpdateProductVersion(product)

	if result != nil {
		return false
	}
	return true
}

func (service userService) GetProductVersionWithId(id int) *dto.ProductVersion {
	return service.repo.GetProductVersionWithId(id)
}

func (service userService) DeleteProductVersion(id int) bool {
	if result := service.repo.DeleteProductInventory(id); result != nil {
		return false
	}

	if result := service.repo.DeleteProductVersion(id); result != nil {
		return false
	}
	return true
}

func (service userService) GetOrderAdminPage() []dto.Order {
	return service.repo.GetOrderAdminPage()
}

func (service userService) GetOrderDetail(id int) []dto.OrderDetail {
	return service.repo.GetOrderDetail(id)
}

func (service userService) GetOrderWithId(id int) *dto.Order {
	return service.repo.GetOrderWithId(id)
}

func (service userService) GetTotalSalesDayNow() *response.Response {
	total_sales, err := service.repo.GetTotalSalesDayNow()

	if err != nil {
		return CreateFailResponse(error_code.GetTotalSalesDayFail, err.Error())
	}

	total_revenue, err := service.repo.GetTotalRevenueDayNow()

	if err != nil {
		return CreateFailResponse(error_code.GetTotalSalesDayFail, err.Error())
	}

	return CreateSuccessResponse(error_code.Success, "", dto.DataSales{
		Sales:   total_sales,
		Revenue: total_revenue,
	})
}

func (service userService) GetTotalSalesWeekNow() *response.Response {
	total_sales, err := service.repo.GetTotalSalesWeekNow()

	if err != nil {
		return CreateFailResponse(error_code.GetTotalSalesDayFail, err.Error())
	}

	total_revenue, err := service.repo.GetTotalRevenueWeekNow()

	if err != nil {
		return CreateFailResponse(error_code.GetTotalSalesDayFail, err.Error())
	}

	return CreateSuccessResponse(error_code.Success, "", dto.DataSales{
		Sales:   total_sales,
		Revenue: total_revenue,
	})
}

func (service userService) GetTotalSalesMonthNow() *response.Response {
	total_sales, err := service.repo.GetTotalSalesMonthNow()

	if err != nil {
		return CreateFailResponse(error_code.GetTotalSalesDayFail, err.Error())
	}

	total_revenue, err := service.repo.GetTotalRevenueMonthNow()

	if err != nil {
		return CreateFailResponse(error_code.GetTotalSalesDayFail, err.Error())
	}

	return CreateSuccessResponse(error_code.Success, "", dto.DataSales{
		Sales:   total_sales,
		Revenue: total_revenue,
	})
}

func (service userService) GetOrdersRecently() *response.Response {
	orders, err := service.repo.GetOrdersRecently()

	if err != nil {
		return CreateFailResponse(error_code.GetOrdersRecentlyFail, err.Error())
	}
	return CreateSuccessResponse(error_code.Success, "", orders)
}

func (service userService) GetTopProducts() *response.Response {
	products, err := service.repo.GetTopProducts()

	if err != nil {
		return CreateFailResponse(error_code.GetTopProductsFail, err.Error())
	}
	return CreateSuccessResponse(error_code.Success, "", products)
}
