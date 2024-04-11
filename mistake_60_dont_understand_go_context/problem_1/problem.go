package main

import (
	"context"
	"fmt"
)

// У ключа типом является пустым интерфейсом потому, что это может привести
// к коллизиям: две функции из разных пакетов могут в качестве ключа
// использовать одно и то же строковое значение. Следовательно, последнее
// из них переопределит предыдущее.

func main() {
	ctx := context.WithValue(context.Background(), "key", "value")

	value := ctx.Value("key")

	fmt.Println(value) // value
}
