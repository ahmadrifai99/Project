package order

import "gorm.io/gorm"

type Repository interface {
	Get() ([]Order, error)
	Create(order Order) (Order, error)
	Update(order Order) (Order, error)
	Delete(ID int) (Order, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Get() ([]Order, error) {
	var orders = []Order{}
	err := r.db.Preload("User").Preload("Items").Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *repository) Create(order Order) (Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) Update(order Order) (Order, error) {
	err := r.db.Save(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *repository) Delete(ID int) (Order, error) {
	var order = Order{}
	err := r.db.Where("id = ?", ID).Find(&order).Error
	if err != nil {
		return order, err
	}

	err = r.db.Delete(&order, ID).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
