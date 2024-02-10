package mysql

import "io"

type Store struct {
}

func (s *Store) StoreCustomer(customer io.Reader) error {
	return nil
}
