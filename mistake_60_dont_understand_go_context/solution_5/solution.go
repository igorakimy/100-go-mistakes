package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Message string

func handler(ctx context.Context, ch chan Message) error {
	for {
		select {
		// Продолжение получения сообщения от ch.
		case msg := <-ch:
			fmt.Println(msg)
			// Какие-то действия с msg.
		// Если контекст выполнен, возврат связанной с ним ошибки.
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	ch := make(chan Message)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go doSomething(ch)

	if err := handler(ctx, ch); err != nil {
		log.Println(err.Error())
	}
}

func doSomething(ch chan Message) {
	for i := 1; i < 6; i++ {
		ch <- Message(strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
