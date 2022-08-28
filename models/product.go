package models

import "time"

type Product struct {
	ID        int       `json:"id" gorm:"PRIMARY_KEY"`
	Title     string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Image     string    `json:"image"  form:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductResponse struct {
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

type ProductTransaction struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (ProductResponse) TableName() string {
	return "products"
}
