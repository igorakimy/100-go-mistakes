package store

import "problem/client"

// Возврат интерфейса считается плохой практикой, т.к.
// это усложняет дизайн и ограничивает гибкость, заставляя
// потребителей использовать один конкретный тип абстракции.

type InMemoryStore struct {
}

func NewInMemoryStore() client.Store {
	return InMemoryStore{}
}

func (i InMemoryStore) Get() {

}

func (i InMemoryStore) Create() {

}
