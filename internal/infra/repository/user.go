package repository

import (
	"errors"
	"log"
	"user-service/internal/core/dto"
	"user-service/internal/core/port/repository"
)

const rowAffected = 1

var (
	rowAffectedError = errors.New("row affected error")
)

type userRepository struct {
	db repository.Database
}

func NewUserRepository(db repository.Database) repository.UserRepository {
	return userRepository{
		db: db,
	}
}

func (repo userRepository) GetUserWithEmail(email string) *dto.UserDTO {
	var result dto.UserDTO
	repo.db.GetDB().Table("Users").Select("email", "password", "first_name", "last_name").Where("email = ?", email).Scan(&result)
	if result.Id == 0 && result.Email == "" {
		return nil
	}
	return &result
}

func (repo userRepository) CreateUser(user *dto.UserDTO) error {
	result := repo.db.GetDB().Create(user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return rowAffectedError
	}

	return nil
}

// func (repo userRepository) CreateUserAddress(user_address *dto.UserAddress) error {
// 	result := repo.db.GetDB().Create(user_address)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	if result.RowsAffected != 1 {
// 		return rowAffectedError
// 	}

// 	return nil
// }

func (repo userRepository) CreateVerificationMail(verify *dto.VerificationEmail) error {
	result := repo.db.GetDB().Create(verify)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected != 1 {
		return rowAffectedError
	}

	return nil
}

func (repo userRepository) GetVerificationWithEmailAndType(email string, type_verify string) *dto.VerificationEmail {
	var result dto.VerificationEmail
	repo.db.GetDB().Table("Email_verification").Where("email = ? AND type = ?", email, type_verify).First(&result)

	if result.Id == 0 {
		return nil
	}

	return &result

}

func (repo userRepository) UpdateVerificationEmail(verify dto.VerificationEmail) error {
	result := repo.db.GetDB().Table("Email_verification").Where("email = ? AND type = ?", verify.Email, verify.Type).Updates(map[string]interface{}{
		"code":      verify.Code,
		"expire_at": verify.Expire_at,
	})
	return result.Error
}

func (repo userRepository) UpdateStatusVerificationEmail(email string, kind string) error {
	result := repo.db.GetDB().Table("Email_verification").Where("email = ? AND type = ?", email, kind).Update("status", true)
	return result.Error
}

// For Admin ----------------------------------------------------

// Get email and password of admin wiht mail primary key
func (repo userRepository) GetAdmin(email string) *dto.AdminDTO {
	var admin *dto.AdminDTO
	result := repo.db.GetDB().Table("admin").First(&admin, "email = ?", email)
	if result.Error != nil {
		return nil
	}

	return admin
}

// Lay tat ca danh muc san pham
func (repo userRepository) GetProductCategories() []dto.ProductCategory {
	var categories []dto.ProductCategory
	repo.db.GetDB().Table("category").Find(&categories)
	return categories
}

func (repo userRepository) GetLastIDCategories() int {
	var last_id int
	repo.db.GetDB().Table("category").Select("id").Order("id desc").Limit(1).Find(&last_id)
	return last_id
}

func (repo userRepository) InsertCategory(category dto.ProductCategory) error {
	result := repo.db.GetDB().Table("category").Select("id", "name", "description", "image").Create(&category)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) DeleteCategory(id int) error {
	result := repo.db.GetDB().Table("category").Where("id = ?", id).Delete(1)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) GetCategoryWithId(id string) *dto.ProductCategory {
	var category *dto.ProductCategory
	repo.db.GetDB().Table("category").Where("id = ?", id).Find(&category)
	return category
}

// func (repo userRepository) UpdateCategory(category *dto.ProductCategory) error {
// 	result := repo.db.GetDB().Table("category").Where("id = ?", category.Id).Updates(
// 		map[string]interface{}{
// 			"name":        category.Name,
// 			"description": category.Description,
// 			"image":       category.Image,
// 		})
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

func (repo userRepository) GetProductForAdmin() []dto.Product {
	var products []dto.Product
	repo.db.GetDB().Table("product").
		Select(
			"product.id as id,product.name as name, product.category_id as category_id, category.name as category_name, product.description as description, product.price as price, product.image").
		Joins("inner join category on category.id = product.category_id").Scan(&products)
	return products
}

func (repo userRepository) InsertProduct(product *dto.Product) error {
	result := repo.db.GetDB().Table("product").
		Select("name", "category_id", "description", "price", "image", "image1", "image2").
		Create(&product)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) UpdateProduct(product *dto.Product) error {
	result := repo.db.GetDB().Table("product").Where("id = ?", product.Id).Updates(
		map[string]interface{}{
			"name":        product.Name,
			"description": product.Description,
			"category_id": product.Category_id,
			"price":       product.Price,
			"image":       product.Image,
			"image1":      product.Image1,
			"image2":      product.Image2,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) DeleteProduct(id int) error {
	result := repo.db.GetDB().Table("product").Where("id = ?", id).Delete(1)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) DeleteAllProductVersion(id int) error {
	result := repo.db.GetDB().Table("product_version").Where("p_id = ?", id).Delete(nil)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) GetProductVersion(id string) []dto.ProductVersion {
	var products []dto.ProductVersion
	repo.db.GetDB().Table("product_version").
		Select(
			"product_version.id as id, product_version.size_product as size_product, product_inventory.quantity as inventory").
		Joins("inner join product_inventory on product_inventory.product_id = product_version.id").
		Where("product_version.p_id = ?", id).Scan(&products)
	return products
}

func (repo userRepository) GetLastIdProduct() int {
	var last_id int
	repo.db.GetDB().Table("product").Select("id").Order("id desc").Limit(1).Find(&last_id)
	return last_id
}

func (repo userRepository) CreateProductVersion(product *dto.ProductVersion) error {
	result := repo.db.GetDB().Table("product_version").Select("id", "p_id", "size_product").Create(&product)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) GetLastIdProductVersion() int {
	var last_id int
	repo.db.GetDB().Table("product_version").Select("id").Order("id desc").Limit(1).Find(&last_id)
	return last_id
}

func (repo userRepository) CreateProductInventory(inventory *dto.ProductInventory) error {
	result := repo.db.GetDB().Table("product_inventory").Select("product_id", "quantity").Create(&inventory)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (repo userRepository) UpdateProductInventory(id int, quantity int) error {
	result := repo.db.GetDB().Table("product_inventory").Where("product_id = ?", id).Updates(
		map[string]interface{}{
			"quantity": quantity,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) UpdateProductVersion(product *dto.ProductVersion) error {
	result := repo.db.GetDB().Table("product_version").Where("id = ?", product.Id).Updates(
		map[string]interface{}{
			"size_product": product.Size_product,
		})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) GetProductVersionWithId(id int) *dto.ProductVersion {
	var product *dto.ProductVersion
	repo.db.GetDB().Table("product_version").
		Select(
			"product_version.id as id,product_version.p_id as p_id, product_version.size_product as size_product, product_inventory.quantity as inventory").
		Joins("inner join product_inventory on product_inventory.product_id = product_version.id").
		Where("product_version.id = ?", id).Find(&product)
	return product
}

func (repo userRepository) DeleteProductVersion(id int) error {
	result := repo.db.GetDB().Table("product_version").Where("id = ?", id).Delete(1)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) DeleteProductInventory(id int) error {
	result := repo.db.GetDB().Table("product_inventory").Where("product_id = ?", id).Delete(1)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo userRepository) GetOrderAdminPage() []dto.Order {
	var orders []dto.Order
	repo.db.GetDB().Table("orders").
		Select("orders.id as id, orders.user_email as user_email, payment_method.name as payment_name, orders.total as total, orders.created_at as created_at, orders.status as status").
		Joins("inner join payment_method on payment_method.id = orders.payment_id").
		Scan(&orders)
	return orders
}

func (repo userRepository) GetOrderDetail(id int) []dto.OrderDetail {
	var orders []dto.OrderDetail
	repo.db.GetDB().Table("order_detail").
		Select("order_detail.id as id, order_detail.order_id as order_id, product.name as product_name, order_detail.quantity as quantity, product_version.size_product as size_product").
		Joins("inner join product_version on product_version.id = order_detail.product_id").
		Joins("inner join product on product_version.p_id = product.id").
		Where("order_detail.order_id = ?", id).
		Scan(&orders)
	log.Println(orders)
	return orders
}

func (repo userRepository) GetOrderWithId(id int) *dto.Order {
	var order *dto.Order
	repo.db.GetDB().Table("orders").
		Select("orders.id as id, orders.user_email as user_email, payment_method.name as payment_name, orders.total as total, orders.created_at as created_at, orders.status as status").
		Joins("inner join payment_method on payment_method.id = orders.payment_id").
		Where("orders.id = ?", id).
		Find(&order)
	return order
}

func (repo userRepository) GetTotalSalesDayNow() (int, error) {
	var total int
	sql := "SELECT COUNT(*) AS total FROM orders WHERE DATE(created_at) = CURDATE()"
	result := repo.db.GetDB().Raw(sql).Scan(&total)

	if result.Error != nil {
		return -1, result.Error
	}
	return total, nil
}

func (repo userRepository) GetTotalRevenueDayNow() (int, error) {
	var total int
	sql := "SELECT SUM(total) AS total FROM orders WHERE DATE(created_at) = CURDATE()"
	result := repo.db.GetDB().Raw(sql).Scan(&total)

	if result.Error != nil {
		return -1, result.Error
	}
	return total, nil
}

func (repo userRepository) GetTotalSalesWeekNow() (int, error) {
	var total int
	sql := "SELECT COUNT(*) AS total FROM orders WHERE WEEK(created_at) = WEEK(CURDATE()) AND YEAR(created_at) = YEAR(CURDATE())"
	result := repo.db.GetDB().Raw(sql).Scan(&total)

	if result.Error != nil {
		return -1, result.Error
	}
	return total, nil
}

func (repo userRepository) GetTotalRevenueWeekNow() (int, error) {
	var total int
	sql := "SELECT SUM(total) AS total FROM orders WHERE WEEK(created_at) = WEEK(CURDATE()) AND YEAR(created_at) = YEAR(CURDATE())"
	result := repo.db.GetDB().Raw(sql).Scan(&total)

	if result.Error != nil {
		return -1, result.Error
	}
	return total, nil
}

func (repo userRepository) GetTotalSalesMonthNow() (int, error) {
	var total int
	sql := "SELECT COUNT(*) AS total FROM orders WHERE MONTH(created_at) = MONTH(CURDATE()) AND YEAR(created_at) = YEAR(CURDATE())"
	result := repo.db.GetDB().Raw(sql).Scan(&total)

	if result.Error != nil {
		return -1, result.Error
	}
	return total, nil
}

func (repo userRepository) GetTotalRevenueMonthNow() (int, error) {
	var total int
	sql := "SELECT SUM(total) AS total FROM orders WHERE MONTH(created_at) = MONTH(CURDATE()) AND YEAR(created_at) = YEAR(CURDATE())"
	result := repo.db.GetDB().Raw(sql).Scan(&total)

	if result.Error != nil {
		return -1, result.Error
	}
	return total, nil
}

func (repo userRepository) GetOrdersRecently() ([]dto.Order, error) {
	var orders []dto.Order
	sql := "SELECT orders.id as id, orders.user_email as email, orders.total as total, ( SELECT SUM(quantity) FROM order_detail WHERE order_detail.order_id = orders.id) as items FROM orders ORDER BY orders.created_at desc LIMIT 10;"
	result := repo.db.GetDB().Raw(sql).Scan(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (repo userRepository) GetTopProducts() ([]dto.Product, error) {
	var products []dto.Product

	sql := "SELECT product.name as name, category.name as category_name, product.price as price  FROM product inner join category on category.id = product.category_id inner join (SELECT product_version.p_id as p_id, SUM(quantity) as quantity FROM product_version inner join order_detail on product_version.id = order_detail.product_id GROUP BY product_version.p_id ORDER BY quantity DESC LIMIT 10) as A on A.p_id = product.id"
	result := repo.db.GetDB().Raw(sql).Scan(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
