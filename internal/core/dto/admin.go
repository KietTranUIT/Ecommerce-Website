package dto

import "time"

type AdminDTO struct {
	Email    string
	Password string
	Name     string
	Position string
	About_me string
}

type ProductCategory struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Product struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Category_id   int    `json:"category_id"`
	Category_name string `json:"category_name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}

type ProductVersion struct {
	Id           int    `json:"id"`
	P_id         int    `json:"p_id"`
	Size_product int    `json:"size_product"`
	Color        string `json:"color"`
	Image        string `json:"image"`
	Inventory    string `json:"inventory"`
}

type ProductInventory struct {
	Id         int `json:"id"`
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
}

type Order struct {
	Id           int        `json:"id"`
	User_email   string     `json:"user_email"`
	Payment_id   int        `json:"payment_id"`
	Payment_name string     `json:"payment_name"`
	Status       string     `json:"status"`
	Total        int        `json:"total"`
	Created_at   *time.Time `json:"created_at"`
}

type OrderDetail struct {
	Id           int    `json:"id"`
	Order_id     int    `json:"order_id"`
	Product_id   int    `json:"product_id"`
	Product_name string `json:"product_name"`
	Quantity     int    `json:"quantity"`
}
