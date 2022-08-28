package toppingsdto

type ToppingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" gorm:"type:varchar(255)" validate:"required"`
	Price int    `json:"price" gorm:"type: int" form:"price" validate:"required"`
	Image string `json:"image" form:"image" gorm:"type:varchar(255)"`
}
