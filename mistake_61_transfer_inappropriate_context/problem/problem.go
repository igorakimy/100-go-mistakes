package main

import (
	"context"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Выполняется некоторая задача по вычислению ответа HTTP.
	response, err := doSomeTask(r.Context(), r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создание горутины для публикации ответа в Kafka.
	go func() {
		_ = publish(r.Context(), response)
		// Какие-то другие действия.
	}()

	// Запись HTTP-ответа.
	writeResponse(response)
}

func doSomeTask(ctx context.Context, r *http.Request) (http.ResponseWriter, error) {
	return nil, nil
}

func publish(ctx context.Context, w http.ResponseWriter) error {
	// Публикация в Kafka.
	return nil
}

func writeResponse(w http.ResponseWriter) {
	_, _ = fmt.Fprint(w, "Some response")
}

func main() {

}
