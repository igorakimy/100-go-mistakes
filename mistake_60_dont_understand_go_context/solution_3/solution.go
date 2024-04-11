package main

import (
	"context"
	"fmt"
)

// Лучшей практикойпри обработке контекстных ключей будет создание
// неэкспортируемого пользовательского типа.

type key string

// Константа myCustomKey не экспортируется. Поэтому нет риска, что другой пакет,
// использующий тот же контекст, может переопределить уже заданное значение.
// Даже если другой пакет создает тот же myCustomKey на основе типа key, это
// будет другой ключ.

const myCustomKey key = "key"

func f(ctx context.Context) {
	ctx = context.WithValue(ctx, myCustomKey, "foo")
	// ...
	fmt.Println(ctx.Value(myCustomKey))
}

func main() {
	f(context.Background())
}
