package request

type SignUpRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	City      string `json:"city"`
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

type DeleteUserAddressRequest struct {
	Id int `json:"id"`
}

type EditUserAddressRequest struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
	City    string `json:"city"`
	Phone   string `json:"phone"`
}

type Detail struct {
	Id       int `json:"product_id"`
	Quantity int `json:"quantity"`
}

// type CreateOrderRequest struct {
// 	User_email string   `json:"user_email"`
// 	Address_id int      `json:"address_id"`
// 	Payment_id int      `json:"payment_id"`
// 	Total      int      `json:"total"`
// 	Products   []Detail `json:"products"`

// }

// Category
type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Person      string `json:"person"`
	Description string `json:"description"`
}

type Item struct {
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type CreateOrderRequest struct {
	User_email string `json:"user_email"`
	Address_id int    `json:"address_id"`
	Payment_id int    `json:"payment_id"`
	Subtotal   int    `json:"subtotal"`
	Cart       []Item `json:"cart"`
}
