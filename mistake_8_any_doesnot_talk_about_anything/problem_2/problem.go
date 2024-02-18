package main

type Customer struct {
	// какой-то код
}

type Contract struct {
	// какой-то код
}

type Store struct{}

func (s *Store) Get(id string) (any, error) {
	// возвращает any
	return nil, nil
}

func (s *Store) Set(id string, v any) error {
	// принимает any
	return nil
}

func main() {
	s := Store{}
	_ = s.Set("foo", 42)
}
