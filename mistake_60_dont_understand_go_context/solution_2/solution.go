package main

import "context"

func CreateFileWatcher(ctx context.Context, filename string) {
	// ...
}

func main() {
	// Создание конекста, который может быть отменен.
	ctx, cancel := context.WithCancel(context.Background())
	// Откладывание вызова cancel.
	defer cancel()

	go func() {
		// Вызов функции с использованием созданного контекста.
		CreateFileWatcher(ctx, "foo.txt")
	}()

	// ...
}
