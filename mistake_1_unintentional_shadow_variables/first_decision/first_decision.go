package main

import (
	"log"
	"net/http"
)

var client *http.Client

func createDefaultClient() (*http.Client, error) {
	return http.DefaultClient, nil
}

func createClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}

func run(tracing bool) error {

	if tracing {
		// Создается временная переменная c
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		// Переменной client присваивается значение
		// этой временой переменной
		client = c
	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		client = c
	}

	// Использование переменной client
	log.Println(client)

	return nil
}

func main() {
	_ = run(true)
}
