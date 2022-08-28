package toppingsdto

type CreateToppingRequest struct {
	Title string `json:"title" form:"title" gorm:"type:varchar(255)" validate:"required"`
	Price int    `json:"price" gorm:"type: int" form:"price" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type:varchar(255)"`
}

type UpdateToppingRequest struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
}
