package main

type Customer struct {
	id string
}

// Создается абстракция хранилища

type customerStorer interface {
	StoreCustomer(Customer) error
}

type CustomerService struct {
	// Отвязывает CustomerService от фактической реализации
	store customerStorer
}

func (cs CustomerService) CreateNewCustomer(id string) error {
	customer := Customer{id: id}
	return cs.store.StoreCustomer(customer)
}
