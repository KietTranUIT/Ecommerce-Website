package request

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
	Code      string `json:"code"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateRequest struct {
	Email string `json:"email"`
	Kind  string `json:"type"`
	Code  string `json:"code"`
}

// Request for admin

type CreateUserAddressRequest struct {
	User_email string `json:"user_email"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
}
