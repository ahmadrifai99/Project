package order

import (
	"assigment2/models/item"
	"assigment2/models/user"
)

type OrderFormatted struct {
	ID           int
	CustomerName string
	OrderedAt    string
	UserId       int
	User         user.UserFormatted
	Items        []item.ItemFormatted
}

func OrderFormatter(order Order) OrderFormatted {
	orderformatted := OrderFormatted{
		ID:           order.ID,
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		UserId:       order.UserId,
		User:         user.UserFormatter(order.User),
		Items:        item.ItemsFormatter(order.Items),
	}

	return orderformatted
}

func OrdersFormatter(orders []Order) []OrderFormatted {
	var ordersFormatted []OrderFormatted
	for _, order := range orders {
		ordersFormatted = append(ordersFormatted, OrderFormatter(order))
	}

	return ordersFormatted
}
