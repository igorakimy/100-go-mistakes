package main

type Customer struct{}

type Contract struct{}

type Store struct{}

func (s *Store) GetContract(id string) (Contract, error) {
	return Contract{}, nil
}

func (s *Store) SetContract(id string, contract Contract) error {
	return nil
}

func (s *Store) GetCustomer(id string) (Customer, error) {
	return Customer{}, nil
}

func (s *Store) SetCustomer(id string, customer Customer) error {
	return nil
}

type ContractStorer interface {
	GetContract(id string) (Contract, error)
	SetContract(id string, contract Contract) error
}
