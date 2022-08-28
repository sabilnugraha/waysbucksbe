package productsdto

type CreateProductRequest struct {
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: int" form:"price" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type UpdateProductRequest struct {
	Title string `json:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" gorm:"type: int" form:"price"`
}
