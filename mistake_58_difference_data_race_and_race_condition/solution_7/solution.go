package main

import "fmt"

func main() {
	i := 0
	// Создается небуферизованный канал.
	ch := make(chan struct{})

	go func() {
		i = 1
		// Чтение из канала.
		<-ch
	}()

	// Запись в канал.
	ch <- struct{}{}

	// Гонка данных отсутствует, т.к. запись гарантировано происходит до чтения.
	fmt.Println(i)
}
