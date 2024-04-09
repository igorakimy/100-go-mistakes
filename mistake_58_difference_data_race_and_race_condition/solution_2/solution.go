package main

import "sync"

// Мьютекс (mutex) - "mutual exclusion" (взаимное исключение), структура,
// которая обеспечивает обращение к т.н. критической секции не более
// одной горутины.

func main() {
	i := 0
	mutex := sync.Mutex{}

	go func() {
		// Начало критического раздела.
		mutex.Lock()
		// Увеличение i на единицу.
		i++
		// Конец критического раздела.
		mutex.Unlock()
	}()

	go func() {
		mutex.Lock()
		i++
		mutex.Unlock()
	}()
}
