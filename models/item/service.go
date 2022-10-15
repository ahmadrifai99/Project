package item

type Service interface {
	Get() ([]Item, error)
	Create(input CreateItem) (Item, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Get() ([]Item, error) {
	items, err := s.repository.Get()
	if err != nil {
		return items, err
	}

	return items, nil
}

func (s *service) Create(input CreateItem) (Item, error) {
	item := Item{
		ItemCode:    input.ItemCode,
		Quantity:    input.Quantity,
		Description: input.Description,
		OrderId:     input.OrderId,
	}

	itemCreated, err := s.repository.Create(item)
	if err != nil {
		return item, err
	}

	return itemCreated, nil
}
