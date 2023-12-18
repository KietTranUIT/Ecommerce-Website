package error_code

type Error_code string

const (
	Signup_success             Error_code = "SIGNUP_SUCCESS"
	Signup_fail                Error_code = "SIGNUP_FAIL"
	Duplicate_code             Error_code = "DUPLICATE ENTRY"
	InternalError_code         Error_code = "INTERNAL_ERROR"
	SendCodeFailError_code     Error_code = "SEND_CODE_FAIL"
	SendCodeSuccess            Error_code = "SEND_CODE_SUCCESS"
	NotAuthenticatedError_code Error_code = "NOT_AUTHENTICATED"
	FailedAuthentication       Error_code = "AUTHENTICATE_FAIL"
	SuccessAuthentication      Error_code = "SUCCESS_AUTHENTICATE"
	ExpiredCode                Error_code = "EXPIRE_CODE"
	LoginError                 Error_code = "LOGIN_FAIL"
	LoginSuccess               Error_code = "LOGIN_SUCCESS"
	Empty                      Error_code = "EMPTY"
)

var (
	DuplicateUserEmail_msg         = "duplicate email user"
	InternalError_msg              = "internal error"
	DuplicateUserTelephone_msg     = "duplicate telephone user"
	SendCodeFailError_msg          = "send code fail"
	SendCodeSuccess_msg            = "send code success"
	DuplicateEmailVerification_msg = "verification exist"
	NotAuthenticated_msg           = "not authenticated"
	FailedAuthentication_msg       = "fail authenticate"
	SuccessAuthentication_msg      = "success authenticate"
	ExpiredCode_msg                = "expire code"
	NotExistUser_msg               = "user does not exist"
	ExistUser_msg                  = "user exist"
	LoginSuccess_msg               = "login success"
	WrongPassword                  = "incorrected password"
	LoginAdminFail_msg1            = "admin not exist"
	LoginAdminFail_msg2            = "password wrong"
)
