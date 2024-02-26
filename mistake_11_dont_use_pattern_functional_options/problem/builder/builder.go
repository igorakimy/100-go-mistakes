package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const defaultHTTPPort = 80

type Config struct {
	Port int
}

// ConfigBuilder структура, содержащая опциональный порт
type ConfigBuilder struct {
	port *int
}

// Port открытый метод для настройки порта
func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

// Build метод для создания структуры config
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	// Основная логика, связанная с управлением портами
	if b.port == nil {
		cfg.Port = defaultHTTPPort
	} else if *b.port == 0 {
		cfg.Port = randomPort()
	} else if *b.port < 0 {
		return Config{}, errors.New("port should be positive")
	} else {
		cfg.Port = *b.port
	}
	return cfg, nil
}

func NewServer(addr string, config Config) (*http.Server, error) {
	return &http.Server{
		Addr: fmt.Sprintf("%s:%d", addr, config.Port),
	}, nil
}

func randomPort() int {
	rand.Seed(time.Now().UnixNano())
	return int(rand.Int63())
}

func main() {
	// Создается строитель Config
	builder := ConfigBuilder{}
	// Задается порт
	builder.Port(8080)
	// Создается структура Config
	cfg, err := builder.Build()
	if err != nil {
		return
	}
	// Передается структура Config
	server, err := NewServer("localhost", cfg)
	if err != nil {
		return
	}
	fmt.Println(server)
}
