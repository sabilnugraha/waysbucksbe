package cartdto

type CreateCart struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	Product_ID int    `json:"product_id"`
	Topping_ID []int  `json:"topping_id"`
	SubTotal   int    `json:"subtotal"`
	Status     string `json:"status"`
}

type UpdateCart struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}

type UpdateCartRequest struct {
	TransactionID int `json:"transaction_id"`
}
