package client

import "solution/store"

type customersGetter interface {
	GetAllCustomers() ([]store.Customer, error)
}
