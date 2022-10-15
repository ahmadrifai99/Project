package item

import "gorm.io/gorm"

type Repository interface {
	Get() ([]Item, error)
	Create(item Item) (Item, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Get() ([]Item, error) {
	var items []Item
	err := r.db.Find(&items).Error
	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *repository) Create(item Item) (Item, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}
