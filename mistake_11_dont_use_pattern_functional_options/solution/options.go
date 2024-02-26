package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	defaultHTTPPort = 80
	defaultTimeout  = 5 * time.Second
)

// Структура конфигурации
type options struct {
	port    *int
	timeout *time.Duration
}

// Option Представление типа функции, который обновляет
// структуру конфигурации
type Option func(options *options) error

// WithPort функция конфигурации, которая обновляет порт
func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(options *options) error {
		options.timeout = &timeout
		return nil
	}
}

// NewServer принимает вариативные аргументы Option
func NewServer(addr string, opts ...Option) (*http.Server, error) {
	// Создается пустая структура options
	var options options
	// Производится перебор всех входных опций
	for _, opt := range opts {
		// Происходит вызов каждой опции, что приводит
		// к изменению общей структуры options
		if err := opt(&options); err != nil {
			return nil, err
		}
	}
	// На этом этапе строится структура options, которая содержит config.
	// Таким образом, мы можем реализовать нашу логику, связанную
	// с конфигурацией порта var port int.
	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	var timeout time.Duration
	if options.timeout == nil {
		timeout = defaultTimeout
	} else {
		timeout = *options.timeout
	}

	// ...

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", addr, port),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}, nil
}

func randomPort() int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Int63())
}

func main() {
	server, err := NewServer("localhost",
		WithPort(8080),
		WithTimeout(time.Second))
	if err != nil {
		return
	}

	fmt.Println(server)
}
