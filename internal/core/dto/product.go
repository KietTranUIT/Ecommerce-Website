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
	Description   string `json:"description"`
	Image1        string `json:"image1"`
	Image2        string `json:"image2"`
}
