package productsdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}
