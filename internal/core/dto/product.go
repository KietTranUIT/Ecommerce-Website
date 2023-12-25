package dto

type ProductDTO struct {
	Id            int    `json:"id"`
	P_id          int    `json:"p_id"`
	Size_product  int    `json:"size_product"`
	Color         string `json:"color"`
	Image         string `json:"image"`
	Name          string `json:"name"`
	Price         int    `json:"price"`
	Category_name string `json:"category_name"`
	Category_id   int    `json:"category_id"`
	Description   string `json:"description"`
	Image1        string `json:"image1"`
	Image2        string `json:"image2"`
}

type Item_value struct {
	Id           int
	Size_product int
}

type Item struct {
	Value         []Item_value
	Name          string
	Price         int
	Image         string
	Image1        string
	Image2        string
	Image3        string
	Description   string
	Category_name string
	Category_id   int
}
