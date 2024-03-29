package main

import "sync"

type InMem struct {
	// Встроенное поле
	sync.Mutex
	m map[string]int
}

func New() *InMem {
	return &InMem{
		m: make(map[string]int),
	}
}

func (i *InMem) Get(key string) (int, bool) {
	i.Lock()
	v, contains := i.m[key]
	i.Unlock()
	return v, contains
}

func main() {
	m := New()
	// Т.к. метод продвигаемый, то будет доступен извне,
	// что является в большинстве случаев нежелательным поведением
	m.Lock() // ??
}
