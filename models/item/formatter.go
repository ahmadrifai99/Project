package item

type ItemFormatted struct {
	ID          int    `json:"id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}

func ItemFormatter(item Item) ItemFormatted {
	itemFormatted := ItemFormatted{
		ID:          item.ID,
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
		OrderId:     item.OrderId,
	}

	return itemFormatted
}

func ItemsFormatter(items []Item) []ItemFormatted {
	var itemsFormatted []ItemFormatted
	for _, item := range items {
		itemsFormatted = append(itemsFormatted, ItemFormatter(item))
	}

	return itemsFormatted
}
