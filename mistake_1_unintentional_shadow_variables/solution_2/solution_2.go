package main

import (
	"log"
	"net/http"
)

var client *http.Client

// Объявляется переменная err
var err error

func createDefaultClient() (*http.Client, error) {
	return http.DefaultClient, nil
}

func createClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}

func run(tracing bool) error {

	if tracing {
		// Используется оператор присваивания,
		// чтобы напрямую присвоить переменной
		// client значение, возвращаемое *http.Client
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}

	// Можно организовать обработку ошибок вне операторов if
	if err != nil {
		return err
	}

	// Использование переменной client
	log.Println(client)

	return nil
}

func main() {
	_ = run(true)
}
