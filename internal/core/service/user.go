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

func (service userService) Login(req request.LoginRequest) *response.Response {
	user := service.repo.GetUserWithEmail(req.Email)

	if user == nil {
		return CreateFailResponse(error_code.LoginError, error_code.NotExistUser_msg)
	}

	if user.Password != req.Password {
		return CreateFailResponse(error_code.LoginError, error_code.WrongPassword)
	}

	return CreateSuccessResponse(error_code.LoginSuccess, error_code.LoginSuccess_msg)
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
	verify := service.repo.GetVerificationWithEmailAndType(req.Email, "sign up")

	if verify.Status == false {
		return CreateFailResponse(error_code.NotAuthenticatedError_code, error_code.NotAuthenticated_msg)
	}

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
	user_address := dto.UserAddress{
		User_id:   user.Id,
		Telephone: req.Telephone,
		Address:   req.Address,
	}

	if err := service.repo.CreateUserAddress(&user_address); err != nil {
		if strings.Contains(err.Error(), duplicateEntry) {
			return CreateFailResponse(error_code.Duplicate_code, error_code.DuplicateUserTelephone_msg)
		}
		return CreateFailResponse(error_code.InternalError_code, error_code.InternalError_msg)
	}

	return CreateSuccessResponse(error_code.Signup_success, "success")
}

func CreateFailResponse(err_code error_code.Error_code, err_msg string, data ...any) *response.Response {
	return &response.Response{
		Status:     false,
		Error_code: err_code,
		Error_msg:  err_msg,
	}
}

func CreateSuccessResponse(err_code error_code.Error_code, err_msg string, data ...any) *response.Response {
	return &response.Response{
		Status:     true,
		Error_code: err_code,
		Error_msg:  err_msg,
	}
}
