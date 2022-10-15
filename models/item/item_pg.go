package item

// import "database/sql"

// const (
// 	addNewItemQuery = `
// 		INSERT INTO items
// 		(
// 			itemCode,
// 			description,
// 			quantity,
// 			orderId
// 		)
// 		VALUES($1, $2, $3, $4)
// 		RETURNING id, item_code, description, quantity, order_id;
// 	`
// )

// type itemPG struct {
// 	db *sql.DB
// }

// func NewItemPG(db *sql.DB) *Repository {
// 	return &itemPG{
// 		db: db,
// 	}
// }

// func (m *itemPG) Create(item Item) (Item, error) {
// 	var Items Item

// 	row := m.db.QueryRow(addNewItemQuery, item.ItemCode, item.Description, item.Quantity, item.OrderId)

// 	err := row.Scan(ID, ItemCode, Description, Quantity, OrderId)

// 	if err != nil {

// 		return nil, err
// 	}

// 	return &Items, nil
// }

// // func (m *itemPG) GetAllItems() ([]*entity.Item, error) {
// // 	var items = []*entity.Item{}

// // 	rows, err := m.db.Query(retrieveAllItemsQuery)

// // 	if err != nil {
// // 		fmt.Println("err querying items", err.Error())
// // 		return nil, err
// // 	}

// // 	for rows.Next() {
// // 		var item entity.Item

// // 		err = rows.Scan(&item.ItemId, &item.ItemCode, &item.Description, &item.Quantity, &item.OrderId)

// // 		if err != nil {
// // 			fmt.Println("err scanning items", err.Error())
// // 			return nil, err
// // 		}

// // 		items = append(items, &item)
// // 	}

// // 	return items, nil
// // }
