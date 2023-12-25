package dto

import "time"

type Order struct {
	Id           int        `json:"id"`
	User_email   string     `json:"user_email"`
	Address_id   int        `json:"address_id"`
	Payment_id   int        `json:"payment_id"`
	Payment_name string     `json:"payment_name"`
	Status       string     `json:"status"`
	Total        int        `json:"total"`
	Items        int        `json:"items"`
	Created_at   *time.Time `json:"created_at"`
}

type OrderDetail struct {
	Id           int    `json:"id"`
	Order_id     int    `json:"order_id"`
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Size_product int    `json:"size_product"`
	Quantity     int    `json:"quantity"`
}
