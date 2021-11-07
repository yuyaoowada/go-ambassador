package models

type Link struct {
	Model
	Code    string    `json:"code"`
	UserId  uint      `json:"user_id"`
	User    User      `json:"user" gorm:"foreignKey:UserId"`
	Product []Product `json:"products" gorm:"many2many:link_products"`
}
