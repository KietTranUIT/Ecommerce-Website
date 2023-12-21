// Data type for Category object
package dto

type Category struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Person      string `json:"person"`
}
