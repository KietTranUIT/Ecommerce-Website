package request

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
}

type AuthenticateRequest struct {
	Email string `json:"email"`
	Kind  string `json:"type"`
	Code  string `json:"code"`
}
