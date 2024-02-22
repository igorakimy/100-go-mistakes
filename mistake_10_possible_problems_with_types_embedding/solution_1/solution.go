package main

import "sync"

type InMem struct {
	// Поскольку мьютекс не встроен и не экспортируется, доступ к нему
	// извне закрыт.
	mu sync.Mutex
	m  map[string]int
}
