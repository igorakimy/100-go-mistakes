package main

import "fmt"

// Закрытие канала происходит до получения замыкания.

func main() {
	i := 0
	ch := make(chan struct{})

	go func() {
		<-ch
		fmt.Println()
	}()

	i++

	// При таком поведении гонки данных не будет.
	close(ch)
}
