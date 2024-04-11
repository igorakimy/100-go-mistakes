package main

import (
	"context"
	"time"
)

// Внутри себя context.WithTimeout создает горутину, которая будет храниться
// в памяти в течение n-времени или до тех пор, пока не будет вызвана функция
// отмены.

type Position struct {
	lat float64
	lon float64
}

type publisher interface {
	Publish(ctx context.Context, position Position) error
}

type publishHandler struct {
	pub publisher
}

func (h publishHandler) publishPosition(position Position) error {
	// Создание контекста, который будет ожидать только 4 секунды.
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	// Откладывание отмены.
	defer cancel()
	// Передача созданного контекста.
	return h.pub.Publish(ctx, position)
}

func main() {

}
