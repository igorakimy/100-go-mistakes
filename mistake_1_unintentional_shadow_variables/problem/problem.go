package main

import (
	"log"
	"net/http"
)

// Объявляется переменная client
var client *http.Client

func createDefaultClient() (*http.Client, error) {
	return http.DefaultClient, nil
}

func createClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}

func run(tracing bool) error {

	if tracing {
		// Создается HTTP-клиент со включенной тассировкой.
		//(Переменная client затенена в этом блоке)
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		// Создается HTTP-клиент по умолчанию.
		//(Переменная client также затенена в этом блоке)
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}

	// Использование переменной client

	return nil
}

func main() {
	_ = run(true)
}
