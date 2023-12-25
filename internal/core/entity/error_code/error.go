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
	Success                    Error_code = "SUCCESS"
	CreateUserAddressFail      Error_code = "CREATE_USER_ADDRESS_FAIL"
	DeleteUserAddressFail      Error_code = "DELETE_USER_ADDRESS_FAIL"
	UpdateUserAddressFail      Error_code = "UPDATE_USER_ADDRESS_FAIL"
	CreateOrderFail            Error_code = "CREATE_ORDER_FAIL"
	DeleteCategoryFail         Error_code = "DELETE_CATEGORY_FAIL"
	DeleteProductFail          Error_code = "DELETE_PRODUCT_FAIL"
	GetAllProductsFail         Error_code = "GET_ALL_PRODUCTS_FAIL"

	// Category Error code
	GetAllCategoryFail Error_code = "GET_ALL_CATEGORY_FAIL"
	CreateCategoryFail Error_code = "CREATE_CATEGORY_FAIL"
	GetCategoryFail    Error_code = "GET_CATEGORY_FAIL"
	UpdateCategoryFail Error_code = "UPDATE_CATEGORY_FAIL"

	// Product
	GetProductsWithCategoryId Error_code = "GET_PRODUCT_FAIL"
	GetTotalSalesDayFail      Error_code = "GET_TOTAL_SALES_DAY_FAIL"
	GetTotalSalesWeekFail     Error_code = "GET_TOTAL_SALES_WEEK_FAIL"
	GetTotalSalesMonthFail    Error_code = "GET_TOTAL_SALES_MONTH_FAIL"
	GetOrdersRecentlyFail     Error_code = "GET_ORDERS_RECENTLY_FAIL"
	GetTopProductsFail        Error_code = "GET_TOP_PRODUCTS_FAIL"
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
