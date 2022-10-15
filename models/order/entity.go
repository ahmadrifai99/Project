package order

import (
	"assigment2/models/item"
	"assigment2/models/user"
)

type Order struct {
	ID           int `gorm:"default:uuid_generate_v3()"`
	CustomerName string
	OrderedAt    string
	UserId       int
	User         user.User
	Items        []item.Item
}
