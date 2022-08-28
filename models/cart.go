package models

import "time"

type Cart struct {
	ID            int       `json:"id" gorm:"PRIMARY_KEY"`
	Product_ID    int       `json:"product_id"`
	TransactionID *int      `json:"transaction_id"`
	Product       Product   `json:"product"`
	ToppingID     []int     `json:"topping_id" gorm:"-"`
	Topping       []Topping `gorm:"many2many:cart_toppings;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID        int       `json:"user_id"`
	SubTotal      int       `json:"subtotal"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"updated_at"`
	Status        string    `json:"status"`
}

type CartResponse struct {
	ID            int                `json:"id"`
	Total         int                `json:"total"`
	TransactionID int                `json:"transaction_id"`
	ProductID     int                `json:"product_id"`
	Product       ProductTransaction `json:"product"`
	ToppingID     []int              `json:"topping_id" gorm:"-"`
	Topping       []Topping          `json:"topping" gorm:"many2many:cart_toppings"`
}
