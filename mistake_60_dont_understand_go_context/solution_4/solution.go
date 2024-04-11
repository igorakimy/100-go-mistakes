package main

import (
	"context"
	"net/http"
)

type key string

// Создание ключа контекста.
const isValidHostKey key = "isValidHost"

func checkValid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверка валидного хоста.
		validHost := r.Host == "acme"
		// Создание нового контекста со значением, чтобы сообщить, действителен
		// ли исходный хост.
		ctx := context.WithValue(r.Context(), isValidHostKey, validHost)
		// Вызов следующего шага с новым контекстом.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {

}
