package order

import "time"

type Service interface {
	Get() ([]Order, error)
	Create(input CreateOrder) (Order, error)
	Update(input UpdateOrder) (Order, error)
	Delete(input DeleteOrder) (Order, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get() ([]Order, error) {
	orders, err := s.repository.Get()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *service) Create(input CreateOrder) (Order, error) {
	order := Order{
		CustomerName: input.CustomerName,
		UserId:       input.UserId,
		OrderedAt:    time.Now().String(),
	}

	createdOrder, err := s.repository.Create(order)
	if err != nil {
		return order, err
	}

	return createdOrder, nil
}

func (s *service) Update(input UpdateOrder) (Order, error) {
	order := Order{
		ID:           input.ID,
		UserId:       input.UserId,
		CustomerName: input.CustomerName,
		OrderedAt:    input.OrderedAt,
	}

	updatedOrder, err := s.repository.Update(order)
	if err != nil {
		return order, err
	}

	return updatedOrder, nil
}

func (s *service) Delete(input DeleteOrder) (Order, error) {
	order, err := s.repository.Delete(input.ID)
	if err != nil {
		return order, err
	}

	return order, nil
}
