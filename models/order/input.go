package order

type CreateOrder struct {
	CustomerName string `json:"customer_name"`
	UserId       int    `json:"user_id"`
}

type UpdateOrder struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	UserId       int    `json:"user_id"`
	OrderedAt    string `json:"ordered_at"`
}

type DeleteOrder struct {
	ID int `uri:"id"`
}
