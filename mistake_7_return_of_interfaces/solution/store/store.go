package store

type InMemoryStore struct {
}

// NewInMemoryStore возвращает конкретную реализацию, а не интерфейс
func NewInMemoryStore() InMemoryStore {
	return InMemoryStore{}
}

func (i InMemoryStore) Get() {

}

func (i InMemoryStore) Create() {

}
